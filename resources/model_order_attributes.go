/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type OrderAttributes struct {
	IsTakeAway    bool      `json:"is_take_away"`
	OrderDate     time.Time `json:"order_date"`
	PaymentMethod int64     `json:"payment_method"`
	TotalPrice    float64   `json:"total_price"`
}
