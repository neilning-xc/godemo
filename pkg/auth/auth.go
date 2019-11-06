package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(str string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(passwordBytes), err
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
