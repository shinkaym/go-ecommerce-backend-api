package repo

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
	panic("unimplement")
}
