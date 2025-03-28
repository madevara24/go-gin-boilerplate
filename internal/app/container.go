package app

import (
	"go-gin-boilerplate/internal/app/repository/user"
	"go-gin-boilerplate/internal/app/usecase/auth/login"
	"go-gin-boilerplate/internal/app/usecase/healthcheck"
	"go-gin-boilerplate/internal/app/usecase/user/register"
	"go-gin-boilerplate/internal/pkg/datasource"
)

type Container struct {
	// PING
	HealthCheckInport healthcheck.Inport

	// USER
	UserRegisterInport register.Inport
	UserLoginInport    login.Inport
}

func NewContainer(datasource *datasource.DataSource) *Container {
	userRepo := user.NewRepo(datasource)
	return &Container{
		// PING
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgre),

		// USER
		UserRegisterInport: register.NewUsecase(userRepo),
		UserLoginInport:    login.NewUsecase(userRepo),
	}
}
