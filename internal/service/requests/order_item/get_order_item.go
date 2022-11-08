package order_item

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetOrderItemRequest struct {
	OrderItemId int64 `url:"-"`
}

func NewGetOrderItemRequest(r *http.Request) (GetOrderItemRequest, error) {
	request := GetOrderItemRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.OrderItemId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
