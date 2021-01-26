package util

import "encoding/base64"

// DecodeBase64 func
func DecodeBase64(v string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(v)
}
