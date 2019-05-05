package users

import "log"

type UserService struct {}

var userRepository UserRepository

func (s UserService) CreateNewData(data interface{}) (interface{}, error) {
	user, err := userRepository.Create(data)
	if err != nil {
		log.Println("error create data, ", err)
		return user, err
	}

	return user, nil
}
