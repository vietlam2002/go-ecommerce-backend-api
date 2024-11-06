package service

import "github.com/vietlam/go-ecommerce-backend-api/internal/repo"

type CostService struct {
	CostRepo *repo.CostRepo
}

func NewCostService() *CostService {
	return &CostService{
		CostRepo: repo.NewCostRepo(),
	}
}

// cs: cost service
func (cs *CostService) GetInfoCost() string {
	return cs.CostRepo.GetInfoCost()
}
