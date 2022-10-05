package order_item

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `create table if not exists order_items
	(
	    order_item_id integer generated always as identity
	        constraint "ORDER_ITEM_pk"
	            primary key,
	    meal_id       integer not null,
	    quantity      integer not null,
	    order_id      integer not null
	        constraint "ORDER_ITEM_orders_null_fk"
	            references orders
	);
	
	alter table order_items
	    owner to postgres;`

	queryDeleteTable = `drop table order_items;`

	queryInsert = `insert into order_items(meal_id, quantity, order_id)
	values ($1, $2, $3) returning order_item_id;`

	querySelect = `select * from order_items;`

	queryUpdate = `update order_items
	set meal_id=$2, quantity=$3, order_id=$4
	where order_item_id=$1;`

	queryDelete = `delete from order_items
	where order_item_id=$1;`

	queryCleanDb = `delete from order_items;`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Insert(orderItem *OrderItem) (int, error) {
	rows, err := repo.db.Queryx(queryInsert, orderItem.MealId, orderItem.Quantity, orderItem.OrderId)
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

func (repo *Repository) Select() ([]OrderItem, error) {
	rows, err := repo.db.Queryx(querySelect)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	orderItem := OrderItem{}
	var orderItemArray []OrderItem
	for rows.Next() {
		err = rows.StructScan(&orderItem)
		if err != nil {
			return nil, err
		}
		orderItemArray = append(orderItemArray, orderItem)
	}
	return orderItemArray, err
}

func (repo *Repository) Delete(id int) error {
	_, err := repo.db.Exec(queryDelete, id)
	return err
}

func (repo *Repository) Update(id int, orderItem *OrderItem) error {
	_, err := repo.db.Queryx(queryUpdate, id, orderItem.MealId, orderItem.Quantity, orderItem.OrderId)
	return err
}

func (repo *Repository) Clean() error {
	_, err := repo.db.Exec(queryCleanDb)
	return err
}
