/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type MealMenu struct {
	Key
	Relationships MealMenuRelationships `json:"relationships"`
}
type MealMenuResponse struct {
	Data     MealMenu `json:"data"`
	Included Included `json:"included"`
}

type MealMenuListResponse struct {
	Data     []MealMenu `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustMealMenu - returns MealMenu from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustMealMenu(key Key) *MealMenu {
	var mealMenu MealMenu
	if c.tryFindEntry(key, &mealMenu) {
		return &mealMenu
	}
	return nil
}
