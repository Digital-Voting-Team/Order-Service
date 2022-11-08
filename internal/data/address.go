package data

import "gitlab.com/distributed_lab/kit/pgdb"

type AddressesQ interface {
	New() AddressesQ

	Get() (*Address, error)
	Select() ([]Address, error)

	Transaction(fn func(q AddressesQ) error) error

	Insert(Address) (Address, error)
	Update(Address) (Address, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) AddressesQ

	FilterById(ids ...int64) AddressesQ
	FilterByBuildingNumber(numbers ...int64) AddressesQ
	FilterByStreet(streets ...string) AddressesQ
	FilterByCity(cities ...string) AddressesQ
	FilterByDistrict(districts ...string) AddressesQ
	FilterByRegion(regions ...string) AddressesQ
	FilterByPostalCode(codes ...string) AddressesQ
}

type Address struct {
	Id          int64  `db:"address_id" structs:"-"`
	BuildingNum int64  `db:"building_number" structs:"building_number"`
	Street      string `db:"street" structs:"street"`
	City        string `db:"city" structs:"city"`
	District    string `db:"district" structs:"district"`
	Region      string `db:"region" structs:"region"`
	PostalCode  string `db:"postal_code" structs:"postal_code"`
}
