package service

import (
	"github.com/shinkaym/go-ecommerce-backend-api/internal/repo"
	"github.com/shinkaym/go-ecommerce-backend-api/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string { return us.userRepo.GetInfoUser() }

// INTERFACE VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	// declare all interfaces here
}

// Resgister implements IUserService.
func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

// Register implements IUserService
func (us *userService) Register(email string, purpose string) int {
	// check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}
