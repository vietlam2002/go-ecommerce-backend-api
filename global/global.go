package global

import (
	"github.com/vietlam/go-ecommerce-backend-api/pkg/logger"
	"github.com/vietlam/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

// Config để truy cập lấy cấu hình mysql, redis,... từ mọi nơi trong project
