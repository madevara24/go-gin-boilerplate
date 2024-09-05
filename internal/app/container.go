package app

import "go-gin-boilerplate/internal/pkg/datasource"

type Container struct {
}

func NewContainer(datasource *datasource.DataSource) *Container {
	return &Container{}
}
