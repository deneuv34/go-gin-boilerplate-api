package users

import "log"

type UserService struct {
}

func (s UserService) CreateNewData(data interface{}) (Users, error) {
	user, err := Create(data)
	if err != nil {
		log.Println("error create data, ", err)
		return user, err
	}

	return user, nil
}
