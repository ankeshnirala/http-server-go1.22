package applications

import "github.com/ankeshnirala/http-server-go1.22/internals/models"

func (s *Storage) HelloWorldHandler1(request interface{}) *models.Data {
	// validate request

	// write actual business logic

	// connect with infrastructure layer if its required
	// maybe if you want to fetch data from database

	// model response and send it as response
	return &models.Data{
		Message: "HelloWorldHandler1",
	}
}

func (s *Storage) HelloWorldHandler2(request interface{}) *models.Data {
	// validate request

	// write actual business logic

	// model response and send it as response
	return &models.Data{
		Message: "HelloWorldHandler2",
	}
}
