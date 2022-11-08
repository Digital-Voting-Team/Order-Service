package delivery

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetDeliveryListRequest struct {
	pgdb.OffsetPageParams
	FilterPriceFrom []float64   `filter:"price_from"`
	FilterPriceTo   []float64   `filter:"price_to"`
	FilterDateFrom  []time.Time `filter:"date_from"`
	FilterDateTo    []time.Time `filter:"date_to"`
	FilterStaffId   []int64     `filter:"staff_id"`
	FilterOrderId   []int64     `filter:"order_id"`
	FilterAddressId []int64     `filter:"address_id"`
}

func NewGetDeliveryListRequest(r *http.Request) (GetDeliveryListRequest, error) {
	var request GetDeliveryListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
