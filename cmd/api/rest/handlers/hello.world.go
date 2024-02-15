package handlers

import (
	"net/http"

	httpsetup "github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/http-setup"
	"github.com/ankeshnirala/http-server-go1.22/internals/applications"
)

func HelloWorldHandler1(w http.ResponseWriter, r *http.Request) {

	result := applications.New().HelloWorldHandler1([]int{})

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
