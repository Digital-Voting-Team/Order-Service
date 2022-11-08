package order

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteOrderRequest struct {
	OrderId int64 `url:"-"`
}

func NewDeleteOrderRequest(r *http.Request) (DeleteOrderRequest, error) {
	request := DeleteOrderRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.OrderId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
