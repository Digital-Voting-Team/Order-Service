package delivery

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

type UpdateDeliveryRequest struct {
	DeliveryId int64 `url:"-" json:"-"`
	Data       resources.Delivery
}

func NewUpdateDeliveryRequest(r *http.Request) (UpdateDeliveryRequest, error) {
	request := UpdateDeliveryRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.DeliveryId = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateDeliveryRequest) validate() error {
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
