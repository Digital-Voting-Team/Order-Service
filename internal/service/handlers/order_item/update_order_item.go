package order_item

import (
	"github.com/spf13/cast"
	"net/http"
	"order-service/internal/data"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/order_item"
	"order-service/resources"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateOrderItemRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	orderItem, err := helpers.OrderItemsQ(r).FilterById(request.OrderItemId).Get()
	if orderItem == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	newOrderItem := data.OrderItem{
		Quantity: request.Data.Attributes.Quantity,
		MealId:   cast.ToInt64(request.Data.Relationships.Meal.Data.ID),
		OrderId:  cast.ToInt64(request.Data.Relationships.Order.Data.ID),
	}

	relateOrder, err := helpers.OrdersQ(r).FilterById(newOrderItem.OrderId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get new order")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	var resultOrderItem data.OrderItem
	resultOrderItem, err = helpers.OrderItemsQ(r).FilterById(orderItem.Id).Update(newOrderItem)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update order item")
		ape.RenderErr(w, problems.InternalError())
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
