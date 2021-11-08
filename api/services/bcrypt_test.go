package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcrypt(t *testing.T) {
	bcrypt := CreateBcryptService(10)

	hash, err := bcrypt.GenerateHashPassword("test")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = bcrypt.CompareHashAndPassword(hash, "test")
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword(hash, "admin")
	assert.Error(t, err)

	bcrypt = CreateBcryptService(32)
	hash, err = bcrypt.GenerateHashPassword("test")
	assert.Error(t, err)
	assert.Empty(t, hash)
}
