package handlers

import (
	"encoding/json"
	"net/http"

	httpsetup "github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/http-setup"
	"github.com/ankeshnirala/http-server-go1.22/internals/applications"
)

func HelloWorldHandler1(w http.ResponseWriter, r *http.Request) {

	var request applications.HelloWorldHandler1Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httpsetup.WriteJSON(w, http.StatusBadRequest, httpsetup.ApiError{Error: "invalid request paramenter"})
		return
	}

	result, err := applications.New().HelloWorldHandler1(request)
	if err != nil {
		httpsetup.WriteJSON(w, http.StatusBadRequest, httpsetup.ApiError{Error: err.Error()})
		return
	}

	response := &httpsetup.StandardResponse{
		Code: http.StatusOK,
		Data: result,
	}

	httpsetup.WriteJSON(w, http.StatusOK, response)
}

func HelloWorldHandler2(w http.ResponseWriter, r *http.Request) {

	result := applications.New().HelloWorldHandler2([]int{})

	response := &httpsetup.StandardResponse{
		Code: http.StatusOK,
		Data: result,
	}

	httpsetup.WriteJSON(w, http.StatusOK, response)
}
