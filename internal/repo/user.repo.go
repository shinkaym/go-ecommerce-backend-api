package repo

import (
	"github.com/shinkaym/go-ecommerce-backend-api/global"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/database"
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

type userRepository struct{
	sqlc *database.Queries
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

// GetUserEmail implements IUserRepository
func (ur *userRepository) GetUserByEmail(email string) bool {
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	// fmt.Println("Email row:", row)
	user, err := ur.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}

	return user.UsrID != NumberNull
}
