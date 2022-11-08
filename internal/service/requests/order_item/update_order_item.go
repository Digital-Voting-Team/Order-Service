package order_item

import (
	"Order-Service/internal/service/helpers"
	"Order-Service/resources"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type UpdateOrderItemRequest struct {
	OrderItemId int64 `url:"-" json:"-"`
	Data        resources.OrderItem
}

func NewUpdateOrderItemRequest(r *http.Request) (UpdateOrderItemRequest, error) {
	request := UpdateOrderItemRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.OrderItemId = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateOrderItemRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/quantity": validation.Validate(&r.Data.Attributes.Quantity, validation.Required,
			validation.By(helpers.IsInteger)),
		"/data/relationships/meal/data/id": validation.Validate(&r.Data.Relationships.Meal.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/order/data/id": validation.Validate(&r.Data.Relationships.Order.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
