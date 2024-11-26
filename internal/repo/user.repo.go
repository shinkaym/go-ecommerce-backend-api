package repo

import (
	"fmt"

	"github.com/shinkaym/go-ecommerce-backend-api/global"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/model"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "User information retrieved successfully"
// }

// INTERFACE VERSION

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

// GetUserEmail implements IUserRepository
func (ur *userRepository) GetUserByEmail(email string) bool {
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	fmt.Println("Email row:", row)
	return row != NumberNull
}
