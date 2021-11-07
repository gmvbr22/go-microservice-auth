package model

type TwoFactorModel interface {
	/*
	 * Verifica se o passcode 2FA é válido
	 */
	ValidatePasscode(passcode, secret string) bool
}
