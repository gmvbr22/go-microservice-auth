package services

import (
	"errors"

	"github.com/gmvbr/captcha/hcaptcha"
)

type HCaptchaResponse struct {
	Success bool `json:"success"`
}

type HCaptchaService struct {
	captcha *hcaptcha.HCaptcha
}

func CreateHCaptchaService(secret string) *HCaptchaService {
	return &HCaptchaService{
		captcha: hcaptcha.NewHCaptcha(secret),
	}
}

func (s *HCaptchaService) ValidateCaptcha(response string) error {
	if response == "" {
		return errors.New("captcha failed")
	}
	result, err := s.captcha.Verify(response)
	if err != nil {
		return err
	}
	if !result.Success {
		return errors.New("captcha failed")
	}
	return nil
}
