package delivery

import (
	"encoding/json"
	"net/http"
	"order-service/internal/service/helpers"
	"order-service/resources"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateDeliveryRequest struct {
	Data resources.Delivery
}

func NewCreateDeliveryRequest(r *http.Request) (CreateDeliveryRequest, error) {
	var request CreateDeliveryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateDeliveryRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/delivery_price": validation.Validate(&r.Data.Attributes.DeliveryPrice, validation.Required,
			validation.By(helpers.IsFloat)),
		"/data/attributes/delivery_date": validation.Validate(&r.Data.Attributes.DeliveryDate, validation.Required,
			validation.By(helpers.IsDate)),
		"/data/relationships/order/data/id": validation.Validate(&r.Data.Relationships.Order.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/address/data/id": validation.Validate(&r.Data.Relationships.Address.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/staff/data/id": validation.Validate(&r.Data.Relationships.Staff.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
