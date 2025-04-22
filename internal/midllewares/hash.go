package middlewares

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// pass + salt = string need to hash
	// Loop hash power(2, 12)
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(byte), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
