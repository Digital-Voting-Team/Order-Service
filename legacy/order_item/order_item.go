package order_item

type OrderItem struct {
	Id       int `db:"order_item_id"`
	MealId   int `db:"meal_id"`
	Quantity int `db:"quantity"`
	OrderId  int `db:"order_id""`
}

func NewOrderItem(mealId int, quantity int, orderId int) *OrderItem {
	return &OrderItem{
		MealId:   mealId,
		Quantity: quantity,
		OrderId:  orderId,
	}
}
