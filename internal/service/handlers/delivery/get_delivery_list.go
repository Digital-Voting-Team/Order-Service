package delivery

import (
	"net/http"
	"order-service/internal/data"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/delivery"
	"order-service/resources"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDeliveryList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetDeliveryListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	deliveriesQ := helpers.DeliveriesQ(r)
	applyFilters(deliveriesQ, request)
	deliveries, err := deliveriesQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get deliveries")
		ape.Render(w, problems.InternalError())
		return
	}
	orders, err := helpers.OrdersQ(r).FilterById(getOrderIds(deliveries)...).Select()
	addresses, err := helpers.AddressesQ(r).FilterById(getAddressIds(deliveries)...).Select()

	response := resources.DeliveryListResponse{
		Data:     newDeliveriesList(deliveries),
		Links:    helpers.GetOffsetLinks(r, request.OffsetPageParams),
		Included: newDeliveryIncluded(orders, addresses),
	}
	ape.Render(w, response)
}

func applyFilters(q data.DeliveriesQ, request requests.GetDeliveryListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterPriceFrom) > 0 {
		q.FilterByPriceFrom(request.FilterPriceFrom...)
	}
	if len(request.FilterPriceTo) > 0 {
		q.FilterByPriceTo(request.FilterPriceTo...)
	}
	if len(request.FilterDateFrom) > 0 {
		q.FilterByDateFrom(request.FilterDateFrom...)
	}
	if len(request.FilterDateTo) > 0 {
		q.FilterByDateTo(request.FilterDateTo...)
	}
	if len(request.FilterOrderId) > 0 {
		q.FilterByOrderId(request.FilterOrderId...)
	}
	if len(request.FilterAddressId) > 0 {
		q.FilterByAddressId(request.FilterAddressId...)
	}
	if len(request.FilterStaffId) > 0 {
		q.FilterByStaffId(request.FilterStaffId...)
	}
}

func newDeliveriesList(deliveries []data.Delivery) []resources.Delivery {
	result := make([]resources.Delivery, len(deliveries))
	for i, delivery := range deliveries {
		result[i] = resources.Delivery{
			Key: resources.NewKeyInt64(delivery.Id, resources.DELIVERY),
			Attributes: resources.DeliveryAttributes{
				DeliveryPrice: delivery.DeliveryPrice,
				DeliveryDate:  delivery.DeliveryDate,
			},
			Relationships: resources.DeliveryRelationships{
				Order: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(delivery.OrderId, 10),
						Type: resources.ORDER,
					},
				},
				Address: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(delivery.AddressId, 10),
						Type: resources.ADDRESS,
					},
				},
				Staff: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(delivery.StaffId, 10),
						Type: resources.STAFF_REF,
					},
				},
			},
		}
	}
	return result
}

func getOrderIds(deliveries []data.Delivery) []int64 {
	orderIDs := make([]int64, len(deliveries))
	for i := 0; i < len(deliveries); i++ {
		orderIDs[i] = deliveries[i].OrderId
	}
	return orderIDs
}

func getAddressIds(deliveries []data.Delivery) []int64 {
	addressIDs := make([]int64, len(deliveries))
	for i := 0; i < len(deliveries); i++ {
		addressIDs[i] = deliveries[i].AddressId
	}
	return addressIDs
}

func newDeliveryIncluded(orders []data.Order, addresses []data.Address) resources.Included {
	result := resources.Included{}
	for _, item := range orders {
		resource := newOrderModel(item)
		result.Add(&resource)
	}
	for _, item := range addresses {
		resource := newAddressModel(item)
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

func newAddressModel(address data.Address) resources.Address {
	return resources.Address{
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
