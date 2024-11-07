package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {

	// *PongController nghĩa là hàm NewPongController sẽ trả về con trỏ đến 1 instance
	// của PongController
	return &PongController{}
	// PongController{} sẽ trả về một instance mới của PongController
	// & để lấy địa chỉ của đối tượng này
}

// uc: user controller
// us: user service
func (p *PongController) Pong(c *gin.Context) {
	fmt.Println("--> My handle")
	name := c.DefaultQuery("name", "vietlam")

	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping" + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "vietlam"},
	})
}
