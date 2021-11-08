package model

type CaptchaModel interface {
	ValidateCaptcha(response string) error
}
