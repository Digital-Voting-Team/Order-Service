package helpers

import (
	"encoding/json"
	cafeResources "github.com/Digital-Voting-Team/cafe-service/resources"
	customerResources "github.com/Digital-Voting-Team/customer-service/resources"
	menuResources "github.com/Digital-Voting-Team/menu-service/resources"
	staffResources "github.com/Digital-Voting-Team/staff-service/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"
)

func ParseMealResponse(r *http.Response) (*menuResources.MealResponse, error) {
	var response menuResources.MealResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal MealResponse")
	}

	return &response, nil
}

func ParseCafeResponse(r *http.Response) (*cafeResources.CafeResponse, error) {
	var response cafeResources.CafeResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal CafeResponse")
	}

	return &response, nil
}

func ParseCustomerResponse(r *http.Response) (*customerResources.CustomerResponse, error) {
	var response customerResources.CustomerResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal CustomerResponse")
	}

	return &response, nil
}

func ParseStaffResponse(r *http.Response) (*staffResources.StaffResponse, error) {
	var response staffResources.StaffResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal StaffResponse")
	}

	return &response, nil
}

func ValidateMeal(token, endpoint string, id int64) (*menuResources.MealResponse, error) {
	req, err := http.NewRequest("GET", endpoint+strconv.FormatInt(id, 10), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build new request")
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to send request, endpoint: "+endpoint)
	}

	return ParseMealResponse(res)
}

func ValidateCafe(token, endpoint string, id int64) (*cafeResources.CafeResponse, error) {
	req, err := http.NewRequest("GET", endpoint+strconv.FormatInt(id, 10), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build new request")
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to send request, endpoint: "+endpoint)
	}

	return ParseCafeResponse(res)
}

func ValidateCustomer(token, endpoint string, id int64) (*customerResources.CustomerResponse, error) {
	req, err := http.NewRequest("GET", endpoint+strconv.FormatInt(id, 10), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build new request")
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to send request, endpoint: "+endpoint)
	}

	return ParseCustomerResponse(res)
}

func ValidateStaff(token, endpoint string, id int64) (*staffResources.StaffResponse, error) {
	req, err := http.NewRequest("GET", endpoint+strconv.FormatInt(id, 10), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build new request")
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to send request, endpoint: "+endpoint)
	}

	return ParseStaffResponse(res)
}
