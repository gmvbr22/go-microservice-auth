package services

import (
	"github.com/pquerna/otp/totp"
)

type OptService struct{}

func CreateOptService() *OptService {
	return &OptService{}
}

func (s *OptService) ValidateToken(passcode, secret string) bool {
	return totp.Validate(passcode, secret)
}
