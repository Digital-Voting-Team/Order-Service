/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Menu struct {
	Key
	Relationships MenuRelationships `json:"relationships"`
}
type MenuResponse struct {
	Data     Menu     `json:"data"`
	Included Included `json:"included"`
}

type MenuListResponse struct {
	Data     []Menu   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustMenu - returns Menu from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustMenu(key Key) *Menu {
	var menu Menu
	if c.tryFindEntry(key, &menu) {
		return &menu
	}
	return nil
}
