package services

import (
	"encoding/json"
	"io"
	"net/http"
)

type GoogleServiceRequest struct {
	RequestURL string
}

type GoogleServiceResponse struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func (s *Storage) GoogleService(request GoogleServiceRequest) (*GoogleServiceResponse, error) {
	// Make the GET request
	resp, err := http.Get(request.RequestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response *GoogleServiceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
