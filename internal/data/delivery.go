package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type DeliveriesQ interface {
	New() DeliveriesQ

	Get() (*Delivery, error)
	Select() ([]Delivery, error)

	Transaction(fn func(q DeliveriesQ) error) error

	Insert(Delivery) (Delivery, error)
	Update(Delivery) (Delivery, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) DeliveriesQ

	FilterById(ids ...int64) DeliveriesQ
	FilterByPriceFrom(prices ...float64) DeliveriesQ
	FilterByPriceTo(prices ...float64) DeliveriesQ
	FilterByDateFrom(dates ...time.Time) DeliveriesQ
	FilterByDateTo(dates ...time.Time) DeliveriesQ
	FilterByOrderId(ids ...int64) DeliveriesQ
	FilterByAddressId(ids ...int64) DeliveriesQ
	FilterByStaffId(ids ...int64) DeliveriesQ

	JoinOrder() DeliveriesQ
	JoinAddress() DeliveriesQ
}

type Delivery struct {
	Id            int64     `db:"delivery_id" structs:"-"`
	OrderId       int64     `db:"order_id" structs:"order_id"`
	AddressId     int64     `db:"address_id" structs:"address_id"`
	StaffId       int64     `db:"staff_id" structs:"staff_id"`
	DeliveryPrice float64   `db:"delivery_price" structs:"delivery_price"`
	DeliveryDate  time.Time `db:"delivery_date" structs:"delivery_date"`
}
