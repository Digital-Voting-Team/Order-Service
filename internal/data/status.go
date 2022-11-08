package data

import "gitlab.com/distributed_lab/kit/pgdb"

type StatusesQ interface {
	New() StatusesQ

	Get() (*Status, error)
	Select() ([]Status, error)

	Transaction(fn func(q StatusesQ) error) error

	Insert(Status) (Status, error)
	Update(Status) (Status, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) StatusesQ

	FilterById(ids ...int64) StatusesQ
	FilterByNames(names ...string) StatusesQ
}

type Status struct {
	Id         int64  `db:"status_id" structs:"-"`
	StatusName string `db:"status_name" structs:"status_name"`
}
