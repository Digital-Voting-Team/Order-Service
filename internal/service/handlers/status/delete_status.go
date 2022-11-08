package status

import (
	"Order-Service/internal/service/helpers"
	requests "Order-Service/internal/service/requests/status"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteStatus(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteStatusRequest(r)
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

	err = helpers.StatusesQ(r).Delete(request.StatusId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete status")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
