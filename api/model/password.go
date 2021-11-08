package model

type PasswordModel interface {
	
	GenerateHashPassword(password string) (string, error)

	CompareHashAndPassword(hashedPassword, password string) error
}
