package auth

import (
	"util"

	"gorm.io/gorm"
)

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
	gorm.Model
}

func init() {
	util.DB.AutoMigrate(&User{})
}

// CreateUser func
func CreateUser(user *User) error {
	tx := util.DB.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindUser func
func FindUser(username string) (User, error) {
	var user User
	tx := util.DB.First(&user, "username = ?", username)

	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

// FindUserByID func
func FindUserByID(userID string) (User, error) {
	var user User
	tx := util.DB.First(&user, userID)

	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}
