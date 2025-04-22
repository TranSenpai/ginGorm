package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "ChuongBeo"

// Hear contains algorithm
// Payload contains claims(email, id, username, role, ...)
// SecretKey: a private string

func GenerateToken(email string, contractID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"contractID": contractID,
		"exp":        time.Now().Add(time.Hour * 2).Unix(),
	})

	str, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return str, nil
}

// func VerifyToken(token string) (int64, error) {
// 	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)

// 		if !ok {
// 			return nil, errors.New("Unexpected signing method")
// 		}

// 		return []byte(secretKey), nil
// 	})

// 	if err != nil {
// 		return 0
// 	}
// }
