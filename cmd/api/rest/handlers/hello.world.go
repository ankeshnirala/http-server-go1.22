package handlers

import (
	"encoding/json"
	"net/http"

	httpsetup "github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/http-setup"
	"github.com/ankeshnirala/http-server-go1.22/internals/applications/models"
	"github.com/ankeshnirala/http-server-go1.22/internals/applications/services"
)

func HelloWorldHandler1(w http.ResponseWriter, r *http.Request) {

	var request *models.HelloWorldHandler1Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httpsetup.WriteJSON(w, http.StatusBadRequest, httpsetup.ApiError{Error: "invalid request paramenter"})
		return
	}

	result, err := services.New().HelloWorldHandler1(request)
	if err != nil {
		httpsetup.WriteJSON(w, http.StatusBadRequest, httpsetup.ApiError{Error: err.Error()})
		return
	}

	response := &httpsetup.DataResponse{
		TotalItems: 10,
		Items:      []interface{}{result},
	}

	httpsetup.WriteJSON(w, http.StatusOK, response)
}

func HelloWorldHandler2(w http.ResponseWriter, r *http.Request) {

	result := services.New().HelloWorldHandler2([]int{})

	response := &httpsetup.DataResponse{
		TotalItems: 10,
		Items:      []interface{}{result},
	}

	httpsetup.WriteJSON(w, http.StatusOK, response)
}
