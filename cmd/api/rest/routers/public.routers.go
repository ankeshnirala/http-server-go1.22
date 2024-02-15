package routers

import (
	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/handlers"
	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/middlewares"
)

func (s *Storage) RegisterPublicRouter() {
	privateRouter := s.routers.Group("/api")

	// add middleware to the router
	privateRouter.Use(middlewares.LoggerMiddleware)

	// Register handlers with the router
	privateRouter.Handle("GET /hello", handlers.HelloWorldHandler2)

}
