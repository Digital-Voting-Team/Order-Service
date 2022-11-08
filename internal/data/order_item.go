package data

import "gitlab.com/distributed_lab/kit/pgdb"

type OrderItemsQ interface {
	New() OrderItemsQ

	Get() (*OrderItem, error)
	Select() ([]OrderItem, error)

	Transaction(fn func(q OrderItemsQ) error) error

	Insert(OrderItem) (OrderItem, error)
	Update(OrderItem) (OrderItem, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) OrderItemsQ

	FilterById(ids ...int64) OrderItemsQ
	FilterByMealId(ids ...int64) OrderItemsQ
	FilterByQuantity(quantities ...int64) OrderItemsQ
	FilterByOrderId(ids ...int64) OrderItemsQ

	JoinOrder() OrderItemsQ
}

type OrderItem struct {
	Id       int64 `db:"order_item_id" structs:"-"`
	MealId   int64 `db:"meal_id" structs:"meal_id"`
	Quantity int64 `db:"quantity" structs:"quantity"`
	OrderId  int64 `db:"order_id" structs:"order_id"`
}
