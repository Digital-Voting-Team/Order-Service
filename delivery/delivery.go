package delivery

import "time"

type Delivery struct {
	Id            int       `db:"delivery_id"`
	OrderId       int       `db:"order_id"`
	AddressId     int       `db:"address_id"`
	StaffId       int       `db:"staff_id"`
	DeliveryPrice float64   `db:"delivery_price"`
	DeliveryDate  time.Time `db:"delivery_date"`
}

func NewDelivery(orderId int, addressId int, staffId int, deliveryPrice float64, deliveryDate time.Time) *Delivery {
	return &Delivery{
		OrderId:       orderId,
		AddressId:     addressId,
		StaffId:       staffId,
		DeliveryPrice: deliveryPrice,
		DeliveryDate:  deliveryDate,
	}
}
