package routers

import (
	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/handlers"
	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/middlewares"
)

func (s *Storage) RegisterPrivateRouter() {
	privateRouter := s.routers.Group("/auth")

	// add middleware to the router
	privateRouter.Use(middlewares.LoggerMiddleware, middlewares.AuthMiddleware)

	// Register handlers with the router
	privateRouter.Handle("POST /hello", handlers.HelloWorldHandler1)

}
