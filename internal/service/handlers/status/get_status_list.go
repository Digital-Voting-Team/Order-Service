package status

import (
	"Order-Service/internal/data"
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/status"
	"Order-Service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetStatusList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetStatusListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	statusesQ := helpers.StatusesQ(r)
	applyFilters(statusesQ, request)
	status, err := statusesQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get status")
		ape.Render(w, problems.InternalError())
		return
	}

	response := resources.StatusListResponse{
		Data:  newStatusesList(status),
		Links: helpers.GetOffsetLinks(r, request.OffsetPageParams),
	}
	ape.Render(w, response)
}

func applyFilters(q data.StatusesQ, request requests.GetStatusListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterStatusName) > 0 {
		q.FilterByNames(request.FilterStatusName...)
	}
}

func newStatusesList(statuses []data.Status) []resources.Status {
	result := make([]resources.Status, len(statuses))
	for i, status := range statuses {
		result[i] = resources.Status{
			Key: resources.NewKeyInt64(status.Id, resources.STATUS),
			Attributes: resources.StatusAttributes{
				StatusName: status.StatusName,
			},
		}
	}
	return result
}
