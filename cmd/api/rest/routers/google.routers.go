package routers

import "github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/handlers"

func (s *Storage) RegisterGoogleRouter() {
	googleRouter := s.routers.Group("/ping")

	googleRouter.Handle("GET /google", handlers.GooglePingHandler)
}
