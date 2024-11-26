package response

const (
	ErrCodeSuccess       = 20001
	ErrCodeParamInvalid  = 20003
	ErrInvalidToken      = 30001
	ErrCodeUserHasExists = 50001
	ErrInvalidOTP        = 30002
	ErrSendEmailOtp      = 30003
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "invalid parameter",
	ErrInvalidToken:      "token is invalid",
	ErrCodeUserHasExists: "user has already registerd",
	ErrInvalidOTP:        "Otp error",
	ErrSendEmailOtp:      "Failed to send email OTP",
}
