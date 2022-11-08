package order

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetOrderListRequest struct {
	pgdb.OffsetPageParams
	FilterDateFrom      []time.Time `filter:"date_from"`
	FilterDateTo        []time.Time `filter:"date_to"`
	FilterPriceFrom     []float64   `filter:"price_from"`
	FilterPriceTo       []float64   `filter:"price_to"`
	FilterCustomerId    []int64     `filter:"customer_id"`
	FilterStaffId       []int64     `filter:"staff_id"`
	FilterPaymentMethod []int64     `filter:"payment_method"`
	FilterIsTakeAway    []bool      `filter:"is_take_away"`
	FilterStatusId      []int64     `filter:"status_id"`
	FilterCafeId        []int64     `filter:"cafe_id"`
}

func NewGetOrderListRequest(r *http.Request) (GetOrderListRequest, error) {
	var request GetOrderListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
