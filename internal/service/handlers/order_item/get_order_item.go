package order_item

import (
	"net/http"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/order_item"
	"order-service/resources"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetOrderItem(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetOrderItemRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resultOrderItem, err := helpers.OrderItemsQ(r).FilterById(request.OrderItemId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get order item from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if resultOrderItem == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	relateOrder, err := helpers.OrdersQ(r).FilterById(resultOrderItem.OrderId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get order")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	var includes resources.Included
	includes.Add(&resources.Order{
		Key: resources.NewKeyInt64(relateOrder.Id, resources.ORDER),
		Attributes: resources.OrderAttributes{
			TotalPrice:    relateOrder.TotalPrice,
			PaymentMethod: relateOrder.PaymentMethod,
			IsTakeAway:    relateOrder.IsTakeAway,
			OrderDate:     relateOrder.OrderDate,
		},
		Relationships: resources.OrderRelationships{
			Status: resources.Relation{
				Data: &resources.Key{
					ID:   strconv.FormatInt(relateOrder.StatusId, 10),
					Type: resources.STATUS,
				},
			},
		},
	})

	result := resources.OrderItemResponse{
		Data: resources.OrderItem{
			Key: resources.NewKeyInt64(resultOrderItem.Id, resources.ORDER_ITEM),
			Attributes: resources.OrderItemAttributes{
				Quantity: resultOrderItem.Quantity,
			},
			Relationships: resources.OrderItemRelationships{
				Meal: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultOrderItem.MealId, 10),
						Type: resources.MEAL_REF,
					},
				},
				Order: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultOrderItem.OrderId, 10),
						Type: resources.ORDER,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
