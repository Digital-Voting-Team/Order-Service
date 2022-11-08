package status

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetStatusListRequest struct {
	pgdb.OffsetPageParams
	FilterStatusName []string `filter:"status_name"`
}

func NewGetStatusListRequest(r *http.Request) (GetStatusListRequest, error) {
	var request GetStatusListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
