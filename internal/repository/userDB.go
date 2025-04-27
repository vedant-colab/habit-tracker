package repository

import (
	"fmt"
	database "habit-tracker/internal/db"
)

type User struct {
	Username string
	Email    string
	Password string
}

type LoginUser struct {
	Username string
	Password string
	ID       uint
}

func GetUserByUsernameOrEmail(username, email string) (User, error) {
	query := "SELECT username, email, password FROM users WHERE username = $1 OR email = $2"
	row := database.DB.QueryRow(query, username, email)

	var user User
	err := row.Scan(&user.Username, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	fmt.Println(user.Email)
	return user, nil
}

func CreateUser(user User) error {
	fmt.Println("This is user:", user)
	query := "insert into users(username, email, password) values ($1, $2, $3)"
	_, err := database.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (LoginUser, error) {
	var user LoginUser
	query := "select id, username, password from users where username = $1"
	row := database.DB.QueryRow(query, username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		fmt.Println("scan err", err)
		return user, err
	}
	return user, nil
}
