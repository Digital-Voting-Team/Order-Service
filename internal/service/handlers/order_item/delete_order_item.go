package order_item

import (
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/order_item"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteOrderItemRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	OrderItem, err := helpers.OrderItemsQ(r).FilterById(request.OrderItemId).Get()
	if OrderItem == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.OrderItemsQ(r).Delete(request.OrderItemId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete order item")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
