package service

import "github.com/vietlam/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserRepo(),
	}
}

// us: user service
func (us *UserService) GetInfoUser() string {
	return us.UserRepo.GetInfoUser()
}
