package delivery

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetDeliveryRequest struct {
	DeliveryId int64 `url:"-"`
}

func NewGetDeliveryRequest(r *http.Request) (GetDeliveryRequest, error) {
	request := GetDeliveryRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.DeliveryId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
