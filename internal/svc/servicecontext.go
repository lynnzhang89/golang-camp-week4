package svc

import "golang-camp-week4/internal/config"

type ServiceContext struct {
}

func NewServiceContext(config config.Config) *ServiceContext {
	return &ServiceContext{}
}
