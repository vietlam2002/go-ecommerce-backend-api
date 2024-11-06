package repo

type CostRepo struct {
}

func NewCostRepo() *CostRepo {
	return &CostRepo{}
}

// cost repo: cr

func (cr *CostRepo) GetInfoCost() string {
	return "100.000vnd"
}
