package delivery

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteDeliveryRequest struct {
	DeliveryId int64 `url:"-"`
}

func NewDeleteDeliveryRequest(r *http.Request) (DeleteDeliveryRequest, error) {
	request := DeleteDeliveryRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.DeliveryId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
