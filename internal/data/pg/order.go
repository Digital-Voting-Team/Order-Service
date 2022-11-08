package pg

import (
	"Order-Service/internal/data"
	"database/sql"
	"fmt"
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const ordersTableName = "public.orders"

func NewOrdersQ(db *pgdb.DB) data.OrdersQ {
	return &ordersQ{
		db:        db.Clone(),
		sql:       sq.Select("orders.*").From(ordersTableName),
		sqlUpdate: sq.Update(ordersTableName).Suffix("returning *"),
	}
}

type ordersQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *ordersQ) New() data.OrdersQ {
	return NewOrdersQ(q.db)
}

func (q *ordersQ) Get() (*data.Order, error) {
	var result data.Order
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *ordersQ) Select() ([]data.Order, error) {
	var result []data.Order
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *ordersQ) Update(order data.Order) (data.Order, error) {
	var result data.Order
	clauses := structs.Map(order)
	clauses["customer_id"] = order.CustomerId
	clauses["staff_id"] = order.StaffId
	clauses["total_price"] = order.TotalPrice
	clauses["payment_method"] = order.PaymentMethod
	clauses["is_take_away"] = order.IsTakeAway
	clauses["status_id"] = order.StatusId
	clauses["cafe_id"] = order.CafeId
	clauses["order_date"] = order.OrderDate

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *ordersQ) Transaction(fn func(q data.OrdersQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *ordersQ) Insert(order data.Order) (data.Order, error) {
	clauses := structs.Map(order)
	clauses["customer_id"] = order.CustomerId
	clauses["staff_id"] = order.StaffId
	clauses["total_price"] = order.TotalPrice
	clauses["payment_method"] = order.PaymentMethod
	clauses["is_take_away"] = order.IsTakeAway
	clauses["status_id"] = order.StatusId
	clauses["cafe_id"] = order.CafeId
	clauses["order_date"] = order.OrderDate

	var result data.Order
	stmt := sq.Insert(ordersTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *ordersQ) Delete(id int64) error {
	stmt := sq.Delete(ordersTableName).Where(sq.Eq{"order_id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *ordersQ) Page(pageParams pgdb.OffsetPageParams) data.OrdersQ {
	q.sql = pageParams.ApplyTo(q.sql, "order_id")
	return q
}

func (q *ordersQ) FilterById(ids ...int64) data.OrdersQ {
	q.sql = q.sql.Where(sq.Eq{"order_id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"order_id": ids})
	return q
}

func (q *ordersQ) FilterByPriceFrom(prices ...float64) data.OrdersQ {
	stmt := sq.GtOrEq{"total_price": prices}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *ordersQ) FilterByPriceTo(prices ...float64) data.OrdersQ {
	stmt := sq.LtOrEq{"total_price": prices}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *ordersQ) FilterByDateFrom(dates ...time.Time) data.OrdersQ {
	stmt := sq.GtOrEq{"order_date": dates}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *ordersQ) FilterByDateTo(dates ...time.Time) data.OrdersQ {
	stmt := sq.LtOrEq{"order_date": dates}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *ordersQ) FilterByCustomerId(ids ...int64) data.OrdersQ {
	q.sql = q.sql.Where(sq.Eq{"customer_id": ids})
	return q
}

func (q *ordersQ) FilterByStaffId(ids ...int64) data.OrdersQ {
	q.sql = q.sql.Where(sq.Eq{"staff_id": ids})
	return q
}

func (q *ordersQ) FilterByPaymentMethod(methods ...int64) data.OrdersQ {
	q.sql = q.sql.Where(sq.Eq{"payment_method": methods})
	return q
}

func (q *ordersQ) FilterByIsTakeAway(b ...bool) data.OrdersQ {
	q.sql = q.sql.Where(sq.Eq{"is_take_away": b})
	return q
}

func (q *ordersQ) FilterByStatusId(ids ...int64) data.OrdersQ {
	q.sql = q.sql.Where(sq.Eq{"status_id": ids})
	return q
}

func (q *ordersQ) FilterByCafeId(ids ...int64) data.OrdersQ {
	q.sql = q.sql.Where(sq.Eq{"cafe_id": ids})
	return q
}

func (q *ordersQ) JoinStatus() data.OrdersQ {
	stmt := fmt.Sprintf("%s as orders on public.statuses.status_id = orders.status_id",
		ordersTableName)
	q.sql = q.sql.Join(stmt)
	return q
}
