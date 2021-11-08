package model

type TokenModel interface {
	GenerateToken(sub, role string, exp int64) (string, error)
}
