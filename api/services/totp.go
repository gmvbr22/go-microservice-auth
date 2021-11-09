package services

import (
	"github.com/pquerna/otp/totp"
)

type TOTPService struct{}

func CreateTOTPService() *TOTPService {
	return &TOTPService{}
}

func (s *TOTPService) ValidatePasscode(passcode, secret string) bool {
	return totp.Validate(passcode, secret)
}
