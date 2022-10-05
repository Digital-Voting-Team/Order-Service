package delivery

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `create table if not exists deliveries
	(
	    delivery_id    integer          not null
	        constraint "DELIVERY_pk"
	            primary key,
	    order_id       integer          not null
	        constraint "DELIVERY_orders_null_fk"
	            references orders,
	    address_id     integer          not null,
	    staff_id       integer          not null,
	    delivery_price double precision not null
	);
	
	alter table deliveries
	    owner to postgres;`

	queryDeleteTable = `drop table deliveries;`

	queryInsert = `insert into deliveries(order_id, address_id, staff_id, delivery_price)
	values ($1) returning delivery_id;`

	querySelect = `select * from deliveries;`

	queryUpdate = `update deliveries
	set order_id=$2, address_id=$3, staff_id=$4, delivery_price=$5
	where delivery_id=$1;`

	queryDelete = `delete from deliveries
	where delivery_id=$1;`

	queryCleanDb = `delete from deliveries;`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Insert(delivery *Delivery) (int, error) {
	rows, err := repo.db.Queryx(queryInsert, delivery.OrderId, delivery.AddressId, delivery.StaffId,
		delivery.DeliveryPrice)
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

func (repo *Repository) Select() ([]Delivery, error) {
	rows, err := repo.db.Queryx(querySelect)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	delivery := Delivery{}
	var deliveryArray []Delivery
	for rows.Next() {
		err = rows.StructScan(&delivery)
		if err != nil {
			return nil, err
		}
		deliveryArray = append(deliveryArray, delivery)
	}
	return deliveryArray, err
}

func (repo *Repository) Delete(id int) error {
	_, err := repo.db.Exec(queryDelete, id)
	return err
}

func (repo *Repository) Update(id int, delivery *Delivery) error {
	_, err := repo.db.Queryx(queryUpdate, id, delivery.OrderId, delivery.AddressId, delivery.StaffId,
		delivery.DeliveryPrice)
	return err
}

func (repo *Repository) Clean() error {
	_, err := repo.db.Exec(queryCleanDb)
	return err
}
