package services

import (
	"github.com/pquerna/otp/totp"
)

type OptService struct{}

func CreateOptService() *OptService {
	return &OptService{}
}

func (s *OptService) ValidatePasscode(passcode, secret string) bool {
	return totp.Validate(passcode, secret)
}
