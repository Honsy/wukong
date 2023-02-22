package userservice

import (
	"test/lib"
	"test/models"
)

type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) GetUserByUsername() (models.User, error) {
	return models.GetUserByUsername(u.Username)
}

func (u *User) InsertUser() error {
	return models.InsertUser(u.Username, lib.MD5(u.Password))
}
