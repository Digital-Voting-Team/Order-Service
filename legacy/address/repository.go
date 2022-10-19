package address

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `create table if not exists addresses
	(
	    address_id   integer generated always as identity
	        constraint addresses_pk
	            primary key,
	    building_number integer not null,
	    street       varchar not null,
	    city         varchar not null,
	    district     varchar not null,
	    region       varchar not null,
	    postal_code  varchar not null
	);
	
	alter table addresses
	    owner to postgres;`

	queryDeleteTable = `drop table addresses;`

	queryInsert = `insert into addresses(building_number, street, city, district, region, postal_code)
	values ($1, $2, $3, $4, $5, $6) returning address_id;`

	querySelect = `select * from addresses;`

	queryUpdate = `update addresses
	set building_number=$2, street=$3, city=$4, district=$5, region=$6, postal_code=$7
	where address_id=$1;`

	queryDelete = `delete from addresses
	where address_id=$1;`

	queryCleanDb = `delete from addresses;`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Insert(address *Address) (int, error) {
	rows, err := repo.db.Queryx(queryInsert, address.BuildingNum, address.Street, address.City, address.District,
		address.Region, address.PostalCode)
	defer rows.Close()
	id := -1
	if err != nil {
		return id, err
	}

	rows.Next()
	err = rows.Scan(&id)
	return id, nil
}

func (repo *Repository) CreateTable() error {
	_, err := repo.db.Exec(queryCreateTable)
	return err
}

func (repo *Repository) DeleteTable() error {
	_, err := repo.db.Exec(queryDeleteTable)
	return err
}

func (repo *Repository) Select() ([]Address, error) {
	rows, err := repo.db.Queryx(querySelect)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	address := Address{}
	var addressArray []Address
	for rows.Next() {
		err = rows.StructScan(&address)
		if err != nil {
			return nil, err
		}
		addressArray = append(addressArray, address)
	}
	return addressArray, err
}

func (repo *Repository) Delete(id int) error {
	_, err := repo.db.Exec(queryDelete, id)
	return err
}

func (repo *Repository) Update(id int, address *Address) error {
	_, err := repo.db.Queryx(queryUpdate, id, address.BuildingNum, address.Street, address.City, address.District,
		address.Region, address.PostalCode)
	return err
}

func (repo *Repository) Clean() error {
	_, err := repo.db.Exec(queryCleanDb)
	return err
}
