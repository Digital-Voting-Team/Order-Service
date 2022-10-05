package status

type Status struct {
	Id         int    `db:"status_id"`
	StatusName string `db:"status_name"`
}

func NewStatus(statusName string) *Status {
	return &Status{StatusName: statusName}
}
