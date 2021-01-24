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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &DecodeToken{
		jwt.StandardClaims{
			Id:        userID,
			NotBefore: time.Now().Unix(),
		},
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		println(err.Error())
	}

	return tokenString
}

// VerifyJwt func
func VerifyJwt(tokenString string, decodeToken *DecodeToken) error {
	token, err := jwt.ParseWithClaims(tokenString, decodeToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		return err
	}
	decodeToken = token.Claims.(*DecodeToken)
	return decodeToken.Valid()
}
