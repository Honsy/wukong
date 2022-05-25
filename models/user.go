package models

import "gorm.io/gorm"

type User struct {
	Model
	Username string `json:"username"`
	Passowrd string `json:"passowrd"`
}

// ExistTagByName checks if there is a tag with the same name
func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.Debug().Select("id").Where("username = ?", user, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return User{}, err
	}

	if user.Username != "" {
		return user, nil
	}

	return User{}, nil
}
