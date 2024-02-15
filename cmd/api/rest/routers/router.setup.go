package routers

import (
	httpsetup "github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/http-setup"
)

type Storage struct {
	routers *httpsetup.Router
}

func New(routers *httpsetup.Router) *Storage {
	return &Storage{
		routers: routers,
	}
}
