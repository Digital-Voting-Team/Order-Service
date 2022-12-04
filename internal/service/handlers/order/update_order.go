package order

import (
	"github.com/spf13/cast"
	"net/http"
	"order-service/internal/data"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/order"
	"order-service/resources"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateOrderRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	order, err := helpers.OrdersQ(r).FilterById(request.OrderId).Get()
	if order == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	newOrder := data.Order{
		TotalPrice:    request.Data.Attributes.TotalPrice,
		PaymentMethod: request.Data.Attributes.PaymentMethod,
		IsTakeAway:    request.Data.Attributes.IsTakeAway,
		OrderDate:     request.Data.Attributes.OrderDate,
		CustomerId:    cast.ToInt64(request.Data.Relationships.Customer.Data.ID),
		StaffId:       cast.ToInt64(request.Data.Relationships.Staff.Data.ID),
		StatusId:      cast.ToInt64(request.Data.Relationships.Status.Data.ID),
		CafeId:        cast.ToInt64(request.Data.Relationships.Cafe.Data.ID),
	}

	relateStatus, err := helpers.StatusesQ(r).FilterById(newOrder.StatusId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get new status")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	var resultOrder data.Order
	resultOrder, err = helpers.OrdersQ(r).FilterById(order.Id).Update(newOrder)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update order")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var includes resources.Included
	includes.Add(&resources.Status{
		Key: resources.NewKeyInt64(relateStatus.Id, resources.STATUS),
		Attributes: resources.StatusAttributes{
			StatusName: relateStatus.StatusName,
		},
	})

	result := resources.OrderResponse{
		Data: resources.Order{
			Key: resources.NewKeyInt64(resultOrder.Id, resources.ORDER),
			Attributes: resources.OrderAttributes{
				TotalPrice:    resultOrder.TotalPrice,
				PaymentMethod: resultOrder.PaymentMethod,
				IsTakeAway:    resultOrder.IsTakeAway,
				OrderDate:     resultOrder.OrderDate,
			},
			Relationships: resources.OrderRelationships{
				Status: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultOrder.StatusId, 10),
						Type: resources.STATUS,
					},
				},
				Customer: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultOrder.CustomerId, 10),
						Type: resources.CUSTOMER_REF,
					},
				},
				Staff: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultOrder.StaffId, 10),
						Type: resources.STAFF_REF,
					},
				},
				Cafe: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultOrder.CafeId, 10),
						Type: resources.CAFE_REF,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
