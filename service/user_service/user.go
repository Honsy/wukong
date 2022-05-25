package userservice

import "test/models"

type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) GetUserByUsername() (models.User, error) {
	return models.GetUserByUsername(u.Username)
}
