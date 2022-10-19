/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type OrderItem struct {
	Key
	Attributes    OrderItemAttributes    `json:"attributes"`
	Relationships OrderItemRelationships `json:"relationships"`
}
type OrderItemResponse struct {
	Data     OrderItem `json:"data"`
	Included Included  `json:"included"`
}

type OrderItemListResponse struct {
	Data     []OrderItem `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustOrderItem - returns OrderItem from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustOrderItem(key Key) *OrderItem {
	var orderItem OrderItem
	if c.tryFindEntry(key, &orderItem) {
		return &orderItem
	}
	return nil
}
