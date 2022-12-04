package address

import (
	"net/http"
	"order-service/internal/data"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/address"
	"order-service/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetAddressList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetAddressListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	addressesQ := helpers.AddressesQ(r)
	applyFilters(addressesQ, request)
	address, err := addressesQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get address")
		ape.Render(w, problems.InternalError())
		return
	}

	response := resources.AddressListResponse{
		Data:  newAddressesList(address),
		Links: helpers.GetOffsetLinks(r, request.OffsetPageParams),
	}
	ape.Render(w, response)
}

func applyFilters(q data.AddressesQ, request requests.GetAddressListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterBuildingNumber) > 0 {
		q.FilterByBuildingNumber(request.FilterBuildingNumber...)
	}

	if len(request.FilterStreet) > 0 {
		q.FilterByStreet(request.FilterStreet...)
	}

	if len(request.FilterCity) > 0 {
		q.FilterByCity(request.FilterCity...)
	}

	if len(request.FilterDistrict) > 0 {
		q.FilterByDistrict(request.FilterDistrict...)
	}

	if len(request.FilterRegion) > 0 {
		q.FilterByRegion(request.FilterRegion...)
	}

	if len(request.FilterPostalCode) > 0 {
		q.FilterByPostalCode(request.FilterPostalCode...)
	}

}

func newAddressesList(addresses []data.Address) []resources.Address {
	result := make([]resources.Address, len(addresses))
	for i, address := range addresses {
		result[i] = resources.Address{
			Key: resources.NewKeyInt64(address.Id, resources.ADDRESS),
			Attributes: resources.AddressAttributes{
				BuildingNumber: address.BuildingNum,
				Street:         address.Street,
				City:           address.City,
				District:       address.District,
				Region:         address.Region,
				PostalCode:     address.PostalCode,
			},
		}
	}
	return result
}
