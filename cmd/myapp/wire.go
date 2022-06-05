//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"golang-camp-week4/internal/svc"
)

// InitializeEvent 声明injector的函数签名
func InitializeServiceContext(configFile string) *svc.ServiceContext {
	wire.Build(NewServiceContext, NewConfig)
	return nil //返回值没有实际意义，只需符合函数签名即可
}
