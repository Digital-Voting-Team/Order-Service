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
