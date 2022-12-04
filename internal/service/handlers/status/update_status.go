package status

import (
	"net/http"
	"order-service/internal/data"
	"order-service/internal/service/helpers"
	requests "order-service/internal/service/requests/status"
	"order-service/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateStatusRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	status, err := helpers.StatusesQ(r).FilterById(request.StatusId).Get()
	if status == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	newStatus := data.Status{
		StatusName: request.Data.Attributes.StatusName,
	}

	var resultStatus data.Status
	resultStatus, err = helpers.StatusesQ(r).FilterById(status.Id).Update(newStatus)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update status")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.StatusResponse{
		Data: resources.Status{
			Key: resources.NewKeyInt64(resultStatus.Id, resources.STATUS),
			Attributes: resources.StatusAttributes{
				StatusName: resultStatus.StatusName,
			},
		},
	}
	ape.Render(w, result)
}
