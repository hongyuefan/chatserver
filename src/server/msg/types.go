package msg

import (
	"errors"
)

type BussTypeId uint32

const (
	Buss_Verify_Code BussTypeId = iota
	Buss_RegistAndLogin_Code
	Buss_GetGameClass_Code
	Buss_Chat_Code
	Buss_Chat_GetFriend_Code
	Buss_Chat_GetBlack_Code
)

var (
	Err_VerificationCode_TimeOut = errors.New("Verification Code TimeOut")
	Err_VerificationCode_Wrong   = errors.New("Verification Code Wrong")
	Err_Token_TimeOut            = errors.New("Token TimeOut")
	Err_Login_NotExist           = errors.New("User Id wrong")
	Err_Get_Friend               = errors.New("Get Friends Error")
	Err_Get_Black                = errors.New("Get Black Error")
)
