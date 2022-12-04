package status

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"order-service/internal/service/helpers"
	"order-service/resources"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type UpdateStatusRequest struct {
	StatusId int64 `url:"-" json:"-"`
	Data     resources.Status
}

func NewUpdateStatusRequest(r *http.Request) (UpdateStatusRequest, error) {
	request := UpdateStatusRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.StatusId = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateStatusRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/status_name": validation.Validate(&r.Data.Attributes.StatusName, validation.Required,
			validation.Length(3, 20)),
	}).Filter()
}
