package rest

import (
	"context"
	"go-gin-boilerplate/internal/app"
	"go-gin-boilerplate/internal/app/delivery/rest/healthcheck"
	"go-gin-boilerplate/internal/pkg/datasource"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router     *gin.Engine
	datasource *datasource.DataSource
	container  *app.Container
}

func NewRouter(ctx context.Context, router *gin.Engine, datasource *datasource.DataSource, container *app.Container) *Router {
	return &Router{
		router:     router,
		datasource: datasource,
		container:  container,
	}
}

func (h *Router) RegisterRouter() {
	h.router.Use(gin.Recovery())

	// PING
	h.router.GET("/health", healthcheck.HealthCheckHandler(h.container.HealthCheckInport))
}
