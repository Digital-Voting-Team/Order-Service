package status

import (
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/status"
	"Order-Service/resources"
	"net/http"

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
