package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/vietlam/go-ecommerce-backend-api/internal/controller"
	"github.com/vietlam/go-ecommerce-backend-api/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}

}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}

}

func CC(c *gin.Context) {

	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")

}

func InitRouter() *gin.Engine {

	r := gin.Default()

	// use the middleware
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)

	v2 := r.Group("v2/2024")
	{
		v2.GET("/ping", c.NewPongController().Pong)
		v2.GET("/user/1", c.NewUserController().GetUserByID) // /v2/2024/ping
		v2.GET("/cost", c.NewCostController().GetCostByID)
		// v2.PUT("/ping", Pong)
	}
	return r
}
