package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type OrdersQ interface {
	New() OrdersQ

	Get() (*Order, error)
	Select() ([]Order, error)

	Transaction(fn func(q OrdersQ) error) error

	Insert(Order) (Order, error)
	Update(Order) (Order, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) OrdersQ

	FilterById(ids ...int64) OrdersQ
	FilterByPriceFrom(prices ...float64) OrdersQ
	FilterByPriceTo(prices ...float64) OrdersQ
	FilterByDateFrom(dates ...time.Time) OrdersQ
	FilterByDateTo(dates ...time.Time) OrdersQ
	FilterByCustomerId(ids ...int64) OrdersQ
	FilterByStaffId(ids ...int64) OrdersQ
	FilterByPaymentMethod(methods ...int64) OrdersQ
	FilterByIsTakeAway(b ...bool) OrdersQ
	FilterByStatusId(ids ...int64) OrdersQ
	FilterByCafeId(ids ...int64) OrdersQ

	JoinStatus() OrdersQ
}

type Order struct {
	Id            int64     `db:"order_id" structs:"-"`
	CustomerId    int64     `db:"customer_id" structs:"customer_id"`
	StaffId       int64     `db:"staff_id" structs:"staff_id"`
	TotalPrice    float64   `db:"total_price" structs:"total_price"`
	PaymentMethod int64     `db:"payment_method" structs:"payment_method"`
	IsTakeAway    bool      `db:"is_take_away" structs:"is_take_away"`
	StatusId      int64     `db:"status_id" structs:"status_id"`
	CafeId        int64     `db:"cafe_id" structs:"cafe_id"`
	OrderDate     time.Time `db:"order_date" structs:"order_date"`
}
