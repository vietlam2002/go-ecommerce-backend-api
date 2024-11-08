package initialize

import (
	"github.com/vietlam/go-ecommerce-backend-api/global"
	"github.com/vietlam/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
