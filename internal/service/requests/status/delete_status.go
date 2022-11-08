package status

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteStatusRequest struct {
	StatusId int64 `url:"-"`
}

func NewDeleteStatusRequest(r *http.Request) (DeleteStatusRequest, error) {
	request := DeleteStatusRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.StatusId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
