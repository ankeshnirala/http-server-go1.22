package routers

import (
	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/handlers"
	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/middlewares"
)

func (s *Storage) RegisterPublicRouter() {
	publicRouter := s.routers.Group("/api")

	// add middleware to the router
	publicRouter.Use(middlewares.LoggerMiddleware)

	// Register handlers with the router
	publicRouter.Handle("GET /hello", handlers.HelloWorldHandler2)

}
