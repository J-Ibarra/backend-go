package auth

import "errors"

// ConfirmPassword func
func ConfirmPassword(user User, password string) error {
	if user.Password == password {
		return nil
	}
	return errors.New("invalid password")
}
