package order_item

import (
	"Order-Service/internal/data"
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/order_item"
	"Order-Service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetOrderItemList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetOrderItemListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	orderItemsQ := helpers.OrderItemsQ(r)
	applyFilters(orderItemsQ, request)
	orderItems, err := orderItemsQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get order items")
		ape.Render(w, problems.InternalError())
		return
	}
	orders, err := helpers.OrdersQ(r).FilterById(getOrderIds(orderItems)...).Select()

	response := resources.OrderItemListResponse{
		Data:     newOrderItemsList(orderItems),
		Links:    helpers.GetOffsetLinks(r, request.OffsetPageParams),
		Included: newOrderItemIncluded(orders),
	}
	ape.Render(w, response)
}

func applyFilters(q data.OrderItemsQ, request requests.GetOrderItemListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterQuantity) > 0 {
		q.FilterByQuantity(request.FilterQuantity...)
	}
	if len(request.FilterMealId) > 0 {
		q.FilterByMealId(request.FilterMealId...)
	}
	if len(request.FilterOrderId) > 0 {
		q.FilterByOrderId(request.FilterOrderId...)
	}
}

func newOrderItemsList(orderItems []data.OrderItem) []resources.OrderItem {
	result := make([]resources.OrderItem, len(orderItems))
	for i, orderItem := range orderItems {
		result[i] = resources.OrderItem{
			Key: resources.NewKeyInt64(orderItem.Id, resources.ORDER_ITEM),
			Attributes: resources.OrderItemAttributes{
				Quantity: orderItem.Quantity,
			},
			Relationships: resources.OrderItemRelationships{
				Meal: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(orderItem.MealId, 10),
						Type: resources.MEAL_REF,
					},
				},
				Order: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(orderItem.OrderId, 10),
						Type: resources.ORDER,
					},
				},
			},
		}
	}
	return result
}

func getOrderIds(orderItems []data.OrderItem) []int64 {
	orderIDs := make([]int64, len(orderItems))
	for i := 0; i < len(orderItems); i++ {
		orderIDs[i] = orderItems[i].OrderId
	}
	return orderIDs
}

func newOrderItemIncluded(orders []data.Order) resources.Included {
	result := resources.Included{}
	for _, item := range orders {
		resource := newOrderModel(item)
		result.Add(&resource)
	}
	return result
}

func newOrderModel(order data.Order) resources.Order {
	return resources.Order{
		Key: resources.NewKeyInt64(order.Id, resources.ORDER),
		Attributes: resources.OrderAttributes{
			TotalPrice:    order.TotalPrice,
			PaymentMethod: order.PaymentMethod,
			IsTakeAway:    order.IsTakeAway,
			OrderDate:     order.OrderDate,
		},
		Relationships: resources.OrderRelationships{
			Status: resources.Relation{
				Data: &resources.Key{
					ID:   strconv.FormatInt(order.StatusId, 10),
					Type: resources.STATUS,
				},
			},
		},
	}
}
