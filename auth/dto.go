package auth

import jwt "github.com/dgrijalva/jwt-go"

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// DecodeToken struct
type DecodeToken struct {
	jwt.StandardClaims
}
