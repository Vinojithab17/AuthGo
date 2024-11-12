package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "Vinojith@AB17"

func GenerateToken(email string, user_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 6).Unix(),
	})
	return token.SignedString([]byte(secretkey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			fmt.Println("invalid token signing")
			return nil, errors.New("invalid token signing")
		}
		return []byte(secretkey), nil
	})

	if err != nil {
		fmt.Println("could not parse token")

		return 0, errors.New("could not parse token")

	}

	isValidToken := parsedToken.Valid

	if !isValidToken {
		fmt.Println("could not parse token")

		return 0, errors.New("invalid token")

	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token signing")
	}
	email := claims["email"].(string)
	user_id := int64(claims["user_id"].(float64))

	fmt.Println(email, user_id)
	return user_id, nil
}
