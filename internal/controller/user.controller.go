package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vietlam/go-ecommerce-backend-api/internal/service"
	"github.com/vietlam/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc: user controller
// us: user service

// CONTROLLER -> SERVICE -> REPO -> MODELS -> DBS
func (uc *UserController) GetUserByID(c *gin.Context) {

	// c.JSON(http.StatusOK, gin.H{ // map string
	// 	// "message": uc.userService.GetInfoUser(),
	// 	// "users":   []string{"cr7", "m10", "vietlam"},

	// 	Code: 1,
	// 	Message: "",
	// 	Data:
	// })

	// if err != nil {
	// 	return response.ErrorResponse(c, 20003, "No need!!")
	// }

	// return response.SuccessResponse(c, 20001, []string{"vietlamhii"})
	response.SuccessResponse(c, 20001, []string{"vietlamhii"})
}
