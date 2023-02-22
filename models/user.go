package models

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// ExistTagByName checks if there is a tag with the same name
func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.Debug().Select("*").Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, err
}

// 插入用户表
func InsertUser(username string, password string) error {
	user := &User{
		Username: username,
		Password: password,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
