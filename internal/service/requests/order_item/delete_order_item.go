package order_item

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteOrderItemRequest struct {
	OrderItemId int64 `url:"-"`
}

func NewDeleteOrderItemRequest(r *http.Request) (DeleteOrderItemRequest, error) {
	request := DeleteOrderItemRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.OrderItemId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
