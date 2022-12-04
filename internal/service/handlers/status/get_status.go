package status

import (
	"net/http"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/status"
	"order-service/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetStatusRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	status, err := helpers.StatusesQ(r).FilterById(request.StatusId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get status from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if status == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	result := resources.StatusResponse{
		Data: resources.Status{
			Key: resources.NewKeyInt64(status.Id, resources.STATUS),
			Attributes: resources.StatusAttributes{
				StatusName: status.StatusName,
			},
		},
	}

	ape.Render(w, result)
}
