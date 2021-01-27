package auth

import "errors"

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userMap map[string]User = make(map[string]User)

func init() {
	userMap["juan"] = User{
		"juan",
		"123456",
	}

	userMap["luis"] = User{
		"luis",
		"123456",
	}

	userMap["manuel"] = User{
		"manuel",
		"123456",
	}
}

// CreateUser func
func CreateUser(user User) {
	userMap[user.Username] = user
}

// FindUser func
func FindUser(username string) (User, error) {
	val, exist := userMap[username]
	if !exist {
		return val, errors.New("user not found")
	}
	return val, nil
}

// ConfirmPassword func
func ConfirmPassword(user User, password string) error {
	if user.Password == password {
		return nil
	}
	return errors.New("invalid password")
}
