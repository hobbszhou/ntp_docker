package svc

import (
	"ntp_server/cmd/api/internal/config"
	"ntp_server/cmd/api/internal/middleware"

	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Cors   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		Cors:   middleware.NewCorsMiddleware().Handle,
	}
}
