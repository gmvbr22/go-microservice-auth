package services

import (
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

func TestTOTP(t *testing.T) {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "gmvbr",
		AccountName: "gmvbr",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	passcode, err := totp.GenerateCode(key.Secret(), time.Now())
	assert.NoError(t, err)
	assert.NotEmpty(t, passcode)

	service := CreateTOTPService()

	ok := service.ValidatePasscode(passcode, key.Secret())
	assert.True(t, ok)

	ok = service.ValidatePasscode("484846", key.Secret())
	assert.False(t, ok)
}
