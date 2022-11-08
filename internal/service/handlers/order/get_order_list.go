package order

import (
	"Order-Service/internal/data"
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/order"
	"Order-Service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetOrderList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetOrderListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ordersQ := helpers.OrdersQ(r)
	applyFilters(ordersQ, request)
	orders, err := ordersQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get orders")
		ape.Render(w, problems.InternalError())
		return
	}
	statuses, err := helpers.StatusesQ(r).FilterById(getStatusIds(orders)...).Select()

	response := resources.OrderListResponse{
		Data:     newOrdersList(orders),
		Links:    helpers.GetOffsetLinks(r, request.OffsetPageParams),
		Included: newOrderIncluded(statuses),
	}
	ape.Render(w, response)
}

func applyFilters(q data.OrdersQ, request requests.GetOrderListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterDateFrom) > 0 {
		q.FilterByDateFrom(request.FilterDateFrom...)
	}
	if len(request.FilterDateTo) > 0 {
		q.FilterByDateTo(request.FilterDateTo...)
	}
	if len(request.FilterPriceFrom) > 0 {
		q.FilterByPriceFrom(request.FilterPriceFrom...)
	}
	if len(request.FilterPriceTo) > 0 {
		q.FilterByPriceTo(request.FilterPriceTo...)
	}
	if len(request.FilterCafeId) > 0 {
		q.FilterByCafeId(request.FilterCafeId...)
	}
	if len(request.FilterCustomerId) > 0 {
		q.FilterByCustomerId(request.FilterCustomerId...)
	}
	if len(request.FilterStaffId) > 0 {
		q.FilterByStaffId(request.FilterStaffId...)
	}
	if len(request.FilterStatusId) > 0 {
		q.FilterByStatusId(request.FilterStatusId...)
	}
	if len(request.FilterPaymentMethod) > 0 {
		q.FilterByPaymentMethod(request.FilterPaymentMethod...)
	}
	if len(request.FilterIsTakeAway) > 0 {
		q.FilterByIsTakeAway(request.FilterIsTakeAway...)
	}
}

func newOrdersList(orders []data.Order) []resources.Order {
	result := make([]resources.Order, len(orders))
	for i, order := range orders {
		result[i] = resources.Order{
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
				Customer: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(order.CustomerId, 10),
						Type: resources.CUSTOMER_REF,
					},
				},
				Staff: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(order.StaffId, 10),
						Type: resources.STAFF_REF,
					},
				},
				Cafe: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(order.CafeId, 10),
						Type: resources.CAFE_REF,
					},
				},
			},
		}
	}
	return result
}

func getStatusIds(orders []data.Order) []int64 {
	statusIDs := make([]int64, len(orders))
	for i := 0; i < len(orders); i++ {
		statusIDs[i] = orders[i].StatusId
	}
	return statusIDs
}

func newOrderIncluded(statuses []data.Status) resources.Included {
	result := resources.Included{}
	for _, item := range statuses {
		resource := newStatusModel(item)
		result.Add(&resource)
	}
	return result
}

func newStatusModel(status data.Status) resources.Status {
	return resources.Status{
		Key: resources.NewKeyInt64(status.Id, resources.STATUS),
		Attributes: resources.StatusAttributes{
			StatusName: status.StatusName,
		},
	}
}
