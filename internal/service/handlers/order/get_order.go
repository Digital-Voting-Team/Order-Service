package order

import (
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/order"
	"Order-Service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetOrderRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resultOrder, err := helpers.OrdersQ(r).FilterById(request.OrderId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get order from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if resultOrder == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	relateStatus, err := helpers.StatusesQ(r).FilterById(resultOrder.StatusId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get status")
		ape.RenderErr(w, problems.NotFound())
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
