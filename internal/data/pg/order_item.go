package pg

import (
	"Order-Service/internal/data"
	"database/sql"
	"fmt"
	"gitlab.com/distributed_lab/kit/pgdb"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const orderItemsTableName = "public.order_items"

func NewOrderItemsQ(db *pgdb.DB) data.OrderItemsQ {
	return &orderItemsQ{
		db:        db.Clone(),
		sql:       sq.Select("orderItems.*").From(orderItemsTableName),
		sqlUpdate: sq.Update(orderItemsTableName).Suffix("returning *"),
	}
}

type orderItemsQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *orderItemsQ) New() data.OrderItemsQ {
	return NewOrderItemsQ(q.db)
}

func (q *orderItemsQ) Get() (*data.OrderItem, error) {
	var result data.OrderItem
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *orderItemsQ) Select() ([]data.OrderItem, error) {
	var result []data.OrderItem
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *orderItemsQ) Update(order data.OrderItem) (data.OrderItem, error) {
	var result data.OrderItem
	clauses := structs.Map(order)
	clauses["meal_id"] = order.MealId
	clauses["quantity"] = order.Quantity
	clauses["order_id"] = order.OrderId

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *orderItemsQ) Transaction(fn func(q data.OrderItemsQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *orderItemsQ) Insert(order data.OrderItem) (data.OrderItem, error) {
	clauses := structs.Map(order)
	clauses["meal_id"] = order.MealId
	clauses["quantity"] = order.Quantity
	clauses["order_id"] = order.OrderId

	var result data.OrderItem
	stmt := sq.Insert(orderItemsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *orderItemsQ) Delete(id int64) error {
	stmt := sq.Delete(orderItemsTableName).Where(sq.Eq{"order_item_id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *orderItemsQ) Page(pageParams pgdb.OffsetPageParams) data.OrderItemsQ {
	q.sql = pageParams.ApplyTo(q.sql, "order_item_id")
	return q
}

func (q *orderItemsQ) FilterById(ids ...int64) data.OrderItemsQ {
	q.sql = q.sql.Where(sq.Eq{"order_item_id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"order_item_id": ids})
	return q
}

func (q *orderItemsQ) FilterByMealId(ids ...int64) data.OrderItemsQ {
	q.sql = q.sql.Where(sq.Eq{"meal_id": ids})
	return q
}

func (q *orderItemsQ) FilterByQuantity(quantities ...int64) data.OrderItemsQ {
	q.sql = q.sql.Where(sq.Eq{"quantity": quantities})
	return q
}

func (q *orderItemsQ) FilterByOrderId(ids ...int64) data.OrderItemsQ {
	q.sql = q.sql.Where(sq.Eq{"order_id": ids})
	return q
}

func (q *orderItemsQ) JoinOrder() data.OrderItemsQ {
	stmt := fmt.Sprintf("%s as order_items on public.orders.order_id = order_items.order_id",
		orderItemsTableName)
	q.sql = q.sql.Join(stmt)
	return q
}
