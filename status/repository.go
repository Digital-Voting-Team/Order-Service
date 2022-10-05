package status

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `create table if not exists statuses
	(
	    status_id   integer generated always as identity
	        constraint "STATUSES_pk"
	            primary key,
	    status_name varchar not null
	);
	
	alter table statuses
	    owner to postgres;`

	queryDeleteTable = `drop table statuses;`

	queryInsert = `insert into statuses(status_name)
	values ($1) returning status_id;`

	querySelect = `select * from statuses;`

	queryUpdate = `update statuses
	set status_name=$2
	where status_id=$1;`

	queryDelete = `delete from statuses
	where status_id=$1;`

	queryCleanDb = `delete from statuses;`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Insert(status *Status) (int, error) {
	rows, err := repo.db.Queryx(queryInsert, status.StatusName)
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

func (repo *Repository) Select() ([]Status, error) {
	rows, err := repo.db.Queryx(querySelect)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	status := Status{}
	var statusArray []Status
	for rows.Next() {
		err = rows.StructScan(&status)
		if err != nil {
			return nil, err
		}
		statusArray = append(statusArray, status)
	}
	return statusArray, err
}

func (repo *Repository) Delete(id int) error {
	_, err := repo.db.Exec(queryDelete, id)
	return err
}

func (repo *Repository) Update(id int, status *Status) error {
	_, err := repo.db.Queryx(queryUpdate, id, status.StatusName)
	return err
}

func (repo *Repository) Clean() error {
	_, err := repo.db.Exec(queryCleanDb)
	return err
}
