package address

import (
	"net/http"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/address"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteAddressRequest(r)
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

	err = helpers.AddressesQ(r).Delete(request.AddressID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete address")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
