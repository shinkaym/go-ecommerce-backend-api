package initialize

import (
	"github.com/shinkaym/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
