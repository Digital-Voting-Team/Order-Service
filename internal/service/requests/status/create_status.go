package status

import (
	"encoding/json"
	"net/http"
	"order-service/internal/service/helpers"
	"order-service/resources"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateStatusRequest struct {
	Data resources.Status
}

func NewCreateStatusRequest(r *http.Request) (CreateStatusRequest, error) {
	var request CreateStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateStatusRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/status_name": validation.Validate(&r.Data.Attributes.StatusName, validation.Required,
			validation.Length(3, 20)),
	}).Filter()
}
