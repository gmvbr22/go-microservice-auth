package services

import (
	"testing"

	"github.com/gmvbr/captcha/hcaptcha_proxy"
	"github.com/stretchr/testify/assert"
)

func handleCaptcha(secret, response string) interface{} {
	result := HCaptchaResponse{}
	result.Success = false

	if response == "20000000-aaaa-bbbb-cccc-000000000001" {
		return nil // Error
	}
	
	if response == "20000000-aaaa-bbbb-cccc-000000000002" {
		result.Success = true
		return result
	}
	return result
}

/*
 * @see test keys, https://docs.hcaptcha.com/#integration-testing-test-keys
 */
func TestHCaptcha(t *testing.T) {
	service := CreateHCaptchaService("0x0000000000000000000000000000000000000000")

	hcaptcha_proxy.Proxy(handleCaptcha, func(url string) {
		service.captcha.UseProxy(url)

		err := service.ValidateCaptcha("")
		assert.Error(t, err)

		err = service.ValidateCaptcha("20000000-aaaa-bbbb-cccc-000000000001")
		assert.Error(t, err)

		err = service.ValidateCaptcha("20000000-aaaa-bbbb-cccc-000000000002")
		assert.NoError(t, err)

		err = service.ValidateCaptcha("20000000-aaaa-bbbb-cccc-000000000003")
		assert.Error(t, err)
	})
}
