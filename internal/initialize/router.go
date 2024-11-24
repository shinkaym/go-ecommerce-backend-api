package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/shinkaym/go-ecommerce-backend-api/global"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {
	// r := gin.Default()
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middleware
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/check_status") // tracking monitor
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	return r
}
