package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/vietlam/go-ecommerce-backend-api/internal/controller"
)

func NewRouter() *gin.Engine {

	r := gin.Default()
	v2 := r.Group("v2/2024")
	{
		v2.GET("/ping", c.NewPongController().Pong)
		v2.GET("/user/1", c.NewUserController().GetUserByID) // /v2/2024/ping
		v2.GET("/cost", c.NewCostController().GetCostByID)
		// v2.PUT("/ping", Pong)
	}
	return r
}
