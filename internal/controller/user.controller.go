package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/service"
	"github.com/shinkaym/go-ecommerce-backend-api/response"
)

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.GetInfoUser())
// }

// INTERFACE VERSION

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
