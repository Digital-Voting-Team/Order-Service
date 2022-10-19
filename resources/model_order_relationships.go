/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type OrderRelationships struct {
	Cafe     Relation                   `json:"cafe"`
	Customer Relation                   `json:"customer"`
	Staff    DeliveryRelationshipsStaff `json:"staff"`
	Status   Relation                   `json:"status"`
}
