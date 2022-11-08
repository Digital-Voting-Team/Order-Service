package delivery

import (
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/delivery"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteDelivery(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteDeliveryRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Delivery, err := helpers.DeliveriesQ(r).FilterById(request.DeliveryId).Get()
	if Delivery == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.DeliveriesQ(r).Delete(request.DeliveryId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete delivery")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
