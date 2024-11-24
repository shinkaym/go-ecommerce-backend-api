package user

import (
	"github.com/gin-gonic/gin"
	// "github.com/shinkaym/go-ecommerce-backend-api/internal/controller"
	// "github.com/shinkaym/go-ecommerce-backend-api/internal/repo"
	// "github.com/shinkaym/go-ecommerce-backend-api/internal/service"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// publish router
	// this is non-dependency
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandlerNonDependency := controller.NewUserController(us)
	userController, _ := wire.InitUserRouterHandler()
	// WIRE go
	// Dependency Injection (DI)
	userRouterPublic := Router.Group("/user")
	{
		// userRouterPublic.POST("/register", userHandlerNonDependency.Register)
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
