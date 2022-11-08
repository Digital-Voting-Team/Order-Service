package status

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetStatusRequest struct {
	StatusId int64 `url:"-"`
}

func NewGetStatusRequest(r *http.Request) (GetStatusRequest, error) {
	request := GetStatusRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.StatusId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
