package response

const (
	ErrCodeSuccess       = 20001
	ErrCodeParamInvalid  = 20003
	ErrInvalidToken      = 30001
	ErrCodeUserHasExists = 50001
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "invalid parameter",
	ErrInvalidToken:      "token is invalid",
	ErrCodeUserHasExists: "user has already registerd",
}
