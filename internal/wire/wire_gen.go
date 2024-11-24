// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/shinkaym/go-ecommerce-backend-api/internal/controller"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/repo"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/service"
)

// Injectors from user.wire.go:

// rename InitUserRouterHandler func to avoid duplication
func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iUserService := service.NewUserService(iUserRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
