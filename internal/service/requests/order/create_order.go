package order

import (
	"encoding/json"
	"net/http"
	"order-service/internal/service/helpers"
	"order-service/resources"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateOrderRequest struct {
	Data resources.Order
}

func NewCreateOrderRequest(r *http.Request) (CreateOrderRequest, error) {
	var request CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateOrderRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/total_price": validation.Validate(&r.Data.Attributes.TotalPrice, validation.Required,
			validation.By(helpers.IsFloat)),
		"/data/attributes/payment_method": validation.Validate(&r.Data.Attributes.PaymentMethod, validation.Required,
			validation.By(helpers.IsInteger)),
		"/data/attributes/order_date": validation.Validate(&r.Data.Attributes.OrderDate, validation.Required,
			validation.By(helpers.IsDate)),
		"/data/relationships/customer/data/id": validation.Validate(&r.Data.Relationships.Customer.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/staff/data/id": validation.Validate(&r.Data.Relationships.Staff.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/status/data/id": validation.Validate(&r.Data.Relationships.Status.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/cafe/data/id": validation.Validate(&r.Data.Relationships.Cafe.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
