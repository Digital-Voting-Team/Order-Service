package order

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `create table if not exists orders
	(
	    order_id       integer generated always as identity
	        constraint "ORDERS_pk"
	            primary key,
	    customer_id    integer          not null,
	    staff_id       integer          not null,
	    total_price    double precision not null,
	    payment_method integer          not null,
	    is_take_away   boolean          not null,
	    status_id      integer          not null
	        constraint "ORDERS_statuses_null_fk"
	            references statuses,
	    cafe_id        integer          not null
	);
	
	alter table orders
	    owner to postgres;`

	queryDeleteTable = `drop table orders;`

	queryInsert = `insert into orders(customer_id, staff_id, total_price, payment_method, is_take_away, status_id, cafe_id)
	values ($1, $2, $3, $4, $5, $6, $7) returning order_id;`

	querySelect = `select * from orders;`

	queryUpdate = `update orders
	set customer_id=$2, staff_id=$3, total_price=$4, payment_method=$5, is_take_away=$6, status_id=$7, cafe_id=$8
	where order_id=$1;`

	queryDelete = `delete from orders
	where order_id=$1;`

	queryCleanDb = `delete from orders;`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Insert(order *Order) (int, error) {
	rows, err := repo.db.Queryx(queryInsert, order.CustomerId, order.StaffId, order.TotalPrice, order.PaymentMethod,
		order.IsTakeAway, order.StatusId, order.CafeId)
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

func (repo *Repository) Select() ([]Order, error) {
	rows, err := repo.db.Queryx(querySelect)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	order := Order{}
	var orderArray []Order
	for rows.Next() {
		err = rows.StructScan(&order)
		if err != nil {
			return nil, err
		}
		orderArray = append(orderArray, order)
	}
	return orderArray, err
}

func (repo *Repository) Delete(id int) error {
	_, err := repo.db.Exec(queryDelete, id)
	return err
}

func (repo *Repository) Update(id int, order *Order) error {
	_, err := repo.db.Queryx(queryUpdate, id, order.CustomerId, order.StaffId, order.TotalPrice, order.PaymentMethod,
		order.IsTakeAway, order.StatusId, order.CafeId)
	return err
}

func (repo *Repository) Clean() error {
	_, err := repo.db.Exec(queryCleanDb)
	return err
}
