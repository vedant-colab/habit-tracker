package service

import (
	"database/sql"
	"fmt"
	"habit-tracker/internal/dto"
	"habit-tracker/internal/repository"
	"habit-tracker/utils"
)

func CreateUserService(userRequest dto.UserRequest) (repository.User, error) {
	hashedPassword, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		return repository.User{}, err
	}
	user := repository.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: hashedPassword,
	}

	// fmt.Println("Hi", userRequest.Username)

	// Check if user already exists
	_, err = repository.GetUserByUsernameOrEmail(user.Username, user.Email)
	if err == nil {
		// Means user already exists
		return repository.User{}, fmt.Errorf("user already exists")
	}

	if err != sql.ErrNoRows {
		// Some unexpected DB error
		return repository.User{}, err
	}

	// Now create the user
	err = repository.CreateUser(user)
	if err != nil {
		return repository.User{}, err
	}

	// Return the created user (You can choose to not return password if you want)
	return user, nil
}

func GetUserByUsername(loginRequest dto.LoginRequest) (repository.LoginUser, error) {
	user := repository.LoginUser{
		Username: loginRequest.Username,
	}
	responseUser, err := repository.GetUserByUsername(user.Username)
	if err != nil {
		return responseUser, err
	}
	return responseUser, nil
}
