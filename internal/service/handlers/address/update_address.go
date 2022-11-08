package address

import (
	"Order-Service/internal/data"
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/address"
	"Order-Service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateAddressRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	address, err := helpers.AddressesQ(r).FilterById(request.AddressID).Get()
	if address == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	newAddress := data.Address{
		BuildingNum: request.Data.Attributes.BuildingNumber,
		Street:      request.Data.Attributes.Street,
		City:        request.Data.Attributes.City,
		District:    request.Data.Attributes.District,
		Region:      request.Data.Attributes.Region,
		PostalCode:  request.Data.Attributes.PostalCode,
	}

	var resultAddress data.Address
	resultAddress, err = helpers.AddressesQ(r).FilterById(address.Id).Update(newAddress)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update address")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.AddressResponse{
		Data: resources.Address{
			Key: resources.NewKeyInt64(resultAddress.Id, resources.ADDRESS),
			Attributes: resources.AddressAttributes{
				BuildingNumber: resultAddress.BuildingNum,
				Street:         resultAddress.Street,
				City:           resultAddress.City,
				District:       resultAddress.District,
				Region:         resultAddress.Region,
				PostalCode:     resultAddress.PostalCode,
			},
		},
	}
	ape.Render(w, result)
}
