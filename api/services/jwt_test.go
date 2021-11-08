package services

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {

	jwt := CreateJWTService([]byte("test"))

	header := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	payload := "eyJleHAiOjUwMDAwLCJyb2xlIjoiZGV2Iiwic3ViIjoiZ212YnIifQ"
	signature := "oCG8ZtJYrhIYu9FdWeRlvHBoHTIxh_56PNiyYfPZdUc"
	expectedToken := strings.Join([]string{header, payload, signature}, ".")

	token, err := jwt.GenerateToken("gmvbr", "dev", 50000)
	assert.NoError(t, err)
	assert.Equal(t, expectedToken, token)

	token, err = jwt.GenerateToken("gmvbr", "dev", 50001)
	assert.NoError(t, err)
	assert.NotEqual(t, expectedToken, token)
}
