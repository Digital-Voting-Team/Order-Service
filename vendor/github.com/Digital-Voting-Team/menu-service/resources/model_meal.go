/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Meal struct {
	Key
	Attributes    MealAttributes    `json:"attributes"`
	Relationships MealRelationships `json:"relationships"`
}
type MealResponse struct {
	Data     Meal     `json:"data"`
	Included Included `json:"included"`
}

type MealListResponse struct {
	Data     []Meal   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustMeal - returns Meal from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustMeal(key Key) *Meal {
	var meal Meal
	if c.tryFindEntry(key, &meal) {
		return &meal
	}
	return nil
}
