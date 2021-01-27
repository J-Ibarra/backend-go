package auth

import (
	"crypto/rsa"
	"time"

	"util"

	jwt "github.com/dgrijalva/jwt-go"
)

type properties struct {
	HmacSecret    string `env:"JWT_HMAC_SECRET,default=ThisItsStrongSecretWordToSigning"`
	RsaPrivateKey string `env:"JWT_RSA_PRIVATE_KEY"`
	RsaPublicKey  string `env:"JWT_RSA_PUBLIC_KEY"`
}

var (
	hmacSecret    []byte
	privateKey    *rsa.PrivateKey
	publicKey     *rsa.PublicKey
	signingMethod jwt.SigningMethod
	isRS256       bool
)

func init() {
	var prop properties
	util.LoadFromEnv(&prop)

	hmacSecret = []byte(prop.HmacSecret)
	signingMethod = jwt.SigningMethodHS256

	if key, err := util.DecodeBase64(prop.RsaPrivateKey); err == nil {
		privateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(key)
	}

	if key, err := util.DecodeBase64(prop.RsaPublicKey); err == nil {
		publicKey, _ = jwt.ParseRSAPublicKeyFromPEM(key)
	}

	if publicKey != nil && privateKey != nil {
		signingMethod = jwt.SigningMethodRS256
		isRS256 = true
	}
}

func generateJwt(userID string) (string, error) {
	token := jwt.NewWithClaims(signingMethod, &DecodeToken{
		jwt.StandardClaims{
			Id:        userID,
			NotBefore: time.Now().Unix(),
		},
	})
	var tokenString string
	var err error

	if isRS256 {
		tokenString, err = token.SignedString(privateKey)
	} else {
		tokenString, err = token.SignedString(hmacSecret)
	}

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJwt func
func VerifyJwt(tokenString string, decodeToken *DecodeToken) error {
	if isRS256 {
		return verify(tokenString, decodeToken, publicKey)
	}
	return verify(tokenString, decodeToken, hmacSecret)
}

func verify(tokenString string, decodeToken *DecodeToken, verifyKey interface{}) error {
	_, err := jwt.ParseWithClaims(tokenString, decodeToken, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		return err
	}
	return decodeToken.Valid()
}
