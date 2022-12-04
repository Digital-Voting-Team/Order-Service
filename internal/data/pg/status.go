package pg

import (
	"database/sql"
	"gitlab.com/distributed_lab/kit/pgdb"
	"order-service/internal/data"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const statusesTableName = "public.statuses"

func NewStatusesQ(db *pgdb.DB) data.StatusesQ {
	return &statusesQ{
		db:        db.Clone(),
		sql:       sq.Select("statuses.*").From(statusesTableName),
		sqlUpdate: sq.Update(statusesTableName).Suffix("returning *"),
	}
}

type statusesQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *statusesQ) New() data.StatusesQ {
	return NewStatusesQ(q.db)
}

func (q *statusesQ) Get() (*data.Status, error) {
	var result data.Status
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *statusesQ) Select() ([]data.Status, error) {
	var result []data.Status
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *statusesQ) Update(status data.Status) (data.Status, error) {
	var result data.Status
	clauses := structs.Map(status)
	clauses["status_name"] = status.StatusName

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *statusesQ) Transaction(fn func(q data.StatusesQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *statusesQ) Insert(status data.Status) (data.Status, error) {
	clauses := structs.Map(status)
	clauses["status_name"] = status.StatusName

	var result data.Status
	stmt := sq.Insert(statusesTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *statusesQ) Delete(id int64) error {
	stmt := sq.Delete(statusesTableName).Where(sq.Eq{"status_id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *statusesQ) Page(pageParams pgdb.OffsetPageParams) data.StatusesQ {
	q.sql = pageParams.ApplyTo(q.sql, "status_id")
	return q
}

func (q *statusesQ) FilterById(ids ...int64) data.StatusesQ {
	q.sql = q.sql.Where(sq.Eq{"status_id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"status_id": ids})
	return q
}

func (q *statusesQ) FilterByNames(names ...string) data.StatusesQ {
	q.sql = q.sql.Where(sq.Eq{"status_name": names})
	return q
}
