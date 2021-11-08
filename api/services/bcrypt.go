package services

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
	cost int
}

func CreateBcryptService(cost int) *BcryptService {
	return &BcryptService{
		cost: cost,
	}
}

func (s *BcryptService) GenerateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *BcryptService) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
