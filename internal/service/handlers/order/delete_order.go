package order

import (
	"net/http"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/order"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteOrderRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Order, err := helpers.OrdersQ(r).FilterById(request.OrderId).Get()
	if Order == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.OrdersQ(r).Delete(request.OrderId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete order")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
