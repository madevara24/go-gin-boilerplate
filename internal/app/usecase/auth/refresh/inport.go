package refresh

import (
	"context"
	"go-gin-boilerplate/internal/app/entity"
)

type Inport interface {
	Execute(ctx context.Context, req InportRequest) (entity.Token, error)
}

type InportRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
