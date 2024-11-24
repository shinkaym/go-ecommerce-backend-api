//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/controller"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/repo"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/service"
)

// rename InitUserRouterHandler func to avoid duplication
func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}