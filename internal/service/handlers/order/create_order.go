package order

import (
	"Order-Service/internal/data"
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/order"
	"Order-Service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateOrderRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Order := data.Order{
		TotalPrice:    request.Data.Attributes.TotalPrice,
		PaymentMethod: request.Data.Attributes.PaymentMethod,
		IsTakeAway:    request.Data.Attributes.IsTakeAway,
		OrderDate:     request.Data.Attributes.OrderDate,
		CustomerId:    cast.ToInt64(request.Data.Relationships.Customer.Data.ID),
		StaffId:       cast.ToInt64(request.Data.Relationships.Staff.Data.ID),
		StatusId:      cast.ToInt64(request.Data.Relationships.Status.Data.ID),
		CafeId:        cast.ToInt64(request.Data.Relationships.Cafe.Data.ID),
	}

	var resultOrder data.Order
	relateStatus, err := helpers.StatusesQ(r).FilterById(Order.StatusId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get status")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resultOrder, err = helpers.OrdersQ(r).Insert(Order)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create order")
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
