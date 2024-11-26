package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shinkaym/go-ecommerce-backend-api/internal/repo"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/utils/random"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/utils/sendto"
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
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
	// declare all interfaces here
}

// Resgister implements IUserService.
func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService
func (us *userService) Register(email string, purpose string) int {
	// 0.hash email
	hashEmail := crypto.GetHash(email)
	fmt.Sprintf("hashEmail::%s", hashEmail)

	// 5.check OTP is available

	// 6.handle user spam

	// 1.check email exists in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2.new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("Otp is ::: %d\n", otp)
	// 3.save OTP in redis with exiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		fmt.Printf("Err ::: %v\n", err)

		return response.ErrInvalidOTP
	}

	// 4.send email OTP
	// make sure the email sent is valid
	err = sendto.SendTemplateEmail([]string{email}, "khoapham1405@gmail.com", "otp-auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})
	// err = sendto.SendTextEmailOtp([]string{email}, os.Getenv("SENDER_EMAIL"), strconv.Itoa(otp))
	if err != nil {
		return response.ErrSendEmailOtp
	}
	return response.ErrCodeSuccess
}
