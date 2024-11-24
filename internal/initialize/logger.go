package initialize

import (
	"github.com/shinkaym/go-ecommerce-backend-api/global"
	"github.com/shinkaym/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
