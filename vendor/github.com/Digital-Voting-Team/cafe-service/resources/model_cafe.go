/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Cafe struct {
	Key
	Attributes    CafeAttributes    `json:"attributes"`
	Relationships CafeRelationships `json:"relationships"`
}
type CafeResponse struct {
	Data     Cafe     `json:"data"`
	Included Included `json:"included"`
}

type CafeListResponse struct {
	Data     []Cafe   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustCafe - returns Cafe from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCafe(key Key) *Cafe {
	var cafe Cafe
	if c.tryFindEntry(key, &cafe) {
		return &cafe
	}
	return nil
}
