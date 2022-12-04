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

func CreateStatus(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateStatusRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var resultStatus data.Status

	status := data.Status{
		StatusName: request.Data.Attributes.StatusName,
	}

	resultStatus, err = helpers.StatusesQ(r).Insert(status)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create status")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.StatusResponse{
		Data: resources.Status{
			Key: resources.NewKeyInt64(resultStatus.Id, resources.STATUS),
			Attributes: resources.StatusAttributes{
				StatusName: request.Data.Attributes.StatusName,
			},
		},
	}
	ape.Render(w, result)
}
