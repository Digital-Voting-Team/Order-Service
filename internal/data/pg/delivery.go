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

const deliveriesTableName = "public.deliveries"

func NewDeliveriesQ(db *pgdb.DB) data.DeliveriesQ {
	return &deliveriesQ{
		db:        db.Clone(),
		sql:       sq.Select("deliveries.*").From(deliveriesTableName),
		sqlUpdate: sq.Update(deliveriesTableName).Suffix("returning *"),
	}
}

type deliveriesQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *deliveriesQ) New() data.DeliveriesQ {
	return NewDeliveriesQ(q.db)
}

func (q *deliveriesQ) Get() (*data.Delivery, error) {
	var result data.Delivery
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *deliveriesQ) Select() ([]data.Delivery, error) {
	var result []data.Delivery
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *deliveriesQ) Update(delivery data.Delivery) (data.Delivery, error) {
	var result data.Delivery
	clauses := structs.Map(delivery)
	clauses["order_id"] = delivery.OrderId
	clauses["address_id"] = delivery.AddressId
	clauses["staff_id"] = delivery.StaffId
	clauses["delivery_price"] = delivery.DeliveryPrice
	clauses["delivery_date"] = delivery.DeliveryDate

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *deliveriesQ) Transaction(fn func(q data.DeliveriesQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *deliveriesQ) Insert(delivery data.Delivery) (data.Delivery, error) {
	clauses := structs.Map(delivery)
	clauses["order_id"] = delivery.OrderId
	clauses["address_id"] = delivery.AddressId
	clauses["staff_id"] = delivery.StaffId
	clauses["delivery_price"] = delivery.DeliveryPrice
	clauses["delivery_date"] = delivery.DeliveryDate

	var result data.Delivery
	stmt := sq.Insert(deliveriesTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *deliveriesQ) Delete(id int64) error {
	stmt := sq.Delete(deliveriesTableName).Where(sq.Eq{"delivery_id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *deliveriesQ) Page(pageParams pgdb.OffsetPageParams) data.DeliveriesQ {
	q.sql = pageParams.ApplyTo(q.sql, "delivery_id")
	return q
}

func (q *deliveriesQ) FilterById(ids ...int64) data.DeliveriesQ {
	q.sql = q.sql.Where(sq.Eq{"delivery_id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"delivery_id": ids})
	return q
}

func (q *deliveriesQ) FilterByPriceFrom(prices ...float64) data.DeliveriesQ {
	stmt := sq.GtOrEq{"delivery_price": prices}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *deliveriesQ) FilterByPriceTo(prices ...float64) data.DeliveriesQ {
	stmt := sq.LtOrEq{"delivery_price": prices}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *deliveriesQ) FilterByDateFrom(dates ...time.Time) data.DeliveriesQ {
	stmt := sq.GtOrEq{"delivery_date": dates}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *deliveriesQ) FilterByDateTo(dates ...time.Time) data.DeliveriesQ {
	stmt := sq.LtOrEq{"delivery_date": dates}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *deliveriesQ) FilterByOrderId(ids ...int64) data.DeliveriesQ {
	q.sql = q.sql.Where(sq.Eq{"order_id": ids})
	return q
}

func (q *deliveriesQ) FilterByAddressId(ids ...int64) data.DeliveriesQ {
	q.sql = q.sql.Where(sq.Eq{"address_id": ids})
	return q
}

func (q *deliveriesQ) FilterByStaffId(ids ...int64) data.DeliveriesQ {
	q.sql = q.sql.Where(sq.Eq{"staff_id": ids})
	return q
}

func (q *deliveriesQ) JoinOrder() data.DeliveriesQ {
	stmt := fmt.Sprintf("%s as deliveries on public.orders.order_id = deliveries.order_id",
		deliveriesTableName)
	q.sql = q.sql.Join(stmt)
	return q
}

func (q *deliveriesQ) JoinAddress() data.DeliveriesQ {
	stmt := fmt.Sprintf("%s as deliveries on public.addresses.address_id = deliveries.address_id",
		deliveriesTableName)
	q.sql = q.sql.Join(stmt)
	return q
}
