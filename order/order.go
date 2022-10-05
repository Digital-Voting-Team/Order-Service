package order

type Order struct {
	Id            int     `db:"order_id"`
	CustomerId    int     `db:"customer_id"`
	StaffId       int     `db:"staff_id"`
	TotalPrice    float64 `db:"total_price"`
	PaymentMethod int     `db:"payment_method"`
	IsTakeAway    bool    `db:"is_take_away"`
	StatusId      int     `db:"status_id"`
	CafeId        int     `db:"cafe_id"`
}

func NewOrder(customerId int, staffId int, totalPrice float64, paymentMethod int, isTakeAway bool, statusId int, cafeId int) *Order {
	return &Order{
		CustomerId:    customerId,
		StaffId:       staffId,
		TotalPrice:    totalPrice,
		PaymentMethod: paymentMethod,
		IsTakeAway:    isTakeAway,
		StatusId:      statusId,
		CafeId:        cafeId,
	}
}
