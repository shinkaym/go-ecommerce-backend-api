package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/service"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/vo"
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
	var params vo.UserRegistratorRequest

	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}

	fmt.Printf("Email param: %s", params.Email)
	
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
