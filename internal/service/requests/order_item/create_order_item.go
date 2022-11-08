package order_item

import (
	"Order-Service/internal/service/helpers"
	"Order-Service/resources"
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateOrderItemRequest struct {
	Data resources.OrderItem
}

func NewCreateOrderItemRequest(r *http.Request) (CreateOrderItemRequest, error) {
	var request CreateOrderItemRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateOrderItemRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/quantity": validation.Validate(&r.Data.Attributes.Quantity, validation.Required,
			validation.By(helpers.IsInteger)),
		"/data/relationships/meal/data/id": validation.Validate(&r.Data.Relationships.Meal.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/order/data/id": validation.Validate(&r.Data.Relationships.Order.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
