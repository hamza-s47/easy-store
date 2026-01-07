package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	const BcryptCost = bcrypt.DefaultCost
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCost)
	return string(bytes), err
}

func ComparePassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
