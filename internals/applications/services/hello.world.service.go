package services

import (
	"github.com/ankeshnirala/http-server-go1.22/internals/applications/models"
	"github.com/ankeshnirala/http-server-go1.22/internals/applications/utils"
)

func (s *Storage) HelloWorldHandler1(request *models.HelloWorldHandler1Request) (*models.Data, error) {
	// validate request
	// var requestData HelloWorldHandler1Request
	// err := json.NewDecoder(request).Decode(&requestData)
	// if err != nil {
	// 	return nil, fmt.Errorf(err)
	// }

	// Create a new instance of Validator
	v := utils.NewValidator()
	// Validate request data
	err := v.ValidateString(request.Email, "Email", "required", "email")
	if err != nil {
		return nil, err
	}

	// write actual business logic

	// connect with infrastructure layer if its required
	// maybe if you want to fetch data from database

	// model response and send it as response
	return &models.Data{
		Message: "HelloWorldHandler1",
	}, nil
}

func (s *Storage) HelloWorldHandler2(request interface{}) *models.Data {
	// validate request

	// write actual business logic

	// model response and send it as response
	return &models.Data{
		Message: "HelloWorldHandler2",
	}
}
