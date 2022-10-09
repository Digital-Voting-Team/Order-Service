package address

type Address struct {
	Id          int    `db:"address_id"`
	BuildingNum int    `db:"building_num"`
	Street      string `db:"street"`
	City        string `db:"city"`
	District    string `db:"district"`
	Region      string `db:"region"`
	PostalCode  string `db:"postal_code"`
}

func NewAddress(buildingNum int, street string, city string, district string, region string, postalCode string) *Address {
	return &Address{
		BuildingNum: buildingNum,
		Street:      street,
		City:        city,
		District:    district,
		Region:      region,
		PostalCode:  postalCode,
	}
}
