package services

import (
	"errors"
	"userauthentication/dto"
	"userauthentication/models"
	"userauthentication/utils"
)

var users = map[string]models.User{}

func Register(userDto dto.UserDto) error {
	var user models.User
	hashedPassword, err := utils.HashPassword(userDto.Password)
	if err != nil {
		return err
	}
	user.Username = userDto.Username
	user.Password = hashedPassword
	users[user.Username] = user
	return nil
}

func Login(userDto dto.UserDto) (models.User, error) {
	user, exists := users[userDto.Username]
	if !exists || !utils.CheckPassword(userDto.Password, user.Password) {
		return models.User{}, errors.New("invalid username or password")
	}
	return user, nil
}

func Profile(username string) (models.User, error) {
	user, ok := users[username]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}