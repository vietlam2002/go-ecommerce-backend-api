package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vietlam/go-ecommerce-backend-api/internal/service"
)

type CostController struct {
	costService *service.CostService
}

func NewCostController() *CostController {
	return &CostController{
		costService: service.NewCostService(),
	}
}

func (cc *CostController) GetCostByID(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": cc.costService.GetInfoCost(),
	})
}
