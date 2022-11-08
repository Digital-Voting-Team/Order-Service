package order

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetOrderRequest struct {
	OrderId int64 `url:"-"`
}

func NewGetOrderRequest(r *http.Request) (GetOrderRequest, error) {
	request := GetOrderRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.OrderId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
