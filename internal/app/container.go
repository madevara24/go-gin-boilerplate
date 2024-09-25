package app

import (
	"go-gin-boilerplate/internal/app/usecase/healthcheck"
	"go-gin-boilerplate/internal/pkg/datasource"
)

type Container struct {
	// PING
	HealthCheckInport healthcheck.Inport
}

func NewContainer(datasource *datasource.DataSource) *Container {
	return &Container{
		// PING
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgre),
	}
}
