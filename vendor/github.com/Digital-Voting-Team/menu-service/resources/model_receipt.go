/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Receipt struct {
	Key
	Attributes    ReceiptAttributes    `json:"attributes"`
	Relationships ReceiptRelationships `json:"relationships"`
}
type ReceiptResponse struct {
	Data     Receipt  `json:"data"`
	Included Included `json:"included"`
}

type ReceiptListResponse struct {
	Data     []Receipt `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustReceipt - returns Receipt from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustReceipt(key Key) *Receipt {
	var receipt Receipt
	if c.tryFindEntry(key, &receipt) {
		return &receipt
	}
	return nil
}
