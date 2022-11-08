package order_item

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetOrderItemListRequest struct {
	pgdb.OffsetPageParams
	FilterMealId   []int64 `filter:"meal_id"`
	FilterQuantity []int64 `filter:"quantity"`
	FilterOrderId  []int64 `filter:"order_id"`
}

func NewGetOrderItemListRequest(r *http.Request) (GetOrderItemListRequest, error) {
	var request GetOrderItemListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
