package model

type TwoFactorModel interface {
	ValidatePasscode(passcode, secret string) bool
}
