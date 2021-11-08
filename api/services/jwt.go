package services

import (
	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	secret []byte
}

func CreateJWTService(secret []byte) *JWTService {
	return &JWTService{secret: secret}
}

func (s *JWTService) GenerateToken(sub, role string, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"exp":  exp,
	})
	return token.SignedString(s.secret)
}
