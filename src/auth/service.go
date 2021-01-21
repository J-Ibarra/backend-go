package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	hmacSampleSecret = []byte("Secret")
)

func generateJwt(userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"nbf": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		println(err.Error())
	}

	claims, err := verifyJwt(tokenString)
	fmt.Printf("%v, %v\n", claims, err)

	return tokenString
}

func verifyJwt(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
