package users

import "fdnBaseAPI/commons/config"

type UserRepository struct{}

func (r UserRepository) CreateNewUser(data interface{}) Users {
	return Users{}
}

func Create(data interface{}) (Users, error) {
	db := config.GetDB()
	var user Users
	err := db.Model(&user).Create(data).Error

	return user, err
}
