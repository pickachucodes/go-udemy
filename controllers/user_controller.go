package controllers

import (
	"encoding/json"
	"mvc/services"
	"mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdParam, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be A num",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		json_value, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(json_value)
		return
	}
	user, apiErr := services.GetUser(userIdParam)
	if apiErr != nil {
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(apiErr.Message))
		return
	}
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
