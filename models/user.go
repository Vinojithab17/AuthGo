package models

import (
	"errors"
	"fmt"
	"time"

	"example.com/go_app/db"
	"example.com/go_app/utils"
)

type User struct {
	ID         int64
	Username   string
	Email      string `binding:"required"`
	Password   string `binding:"required"`
	Created_at time.Time
}

func (u User) AddUser() (int64, error) {
	//later: store to database
	query := `INSERT INTO users(username,email,password,created_at)
				VALUES(?,?,?,?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error Happaned in Preparing the query")
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(u.Username, u.Email, u.Password, u.Created_at)
	if err != nil {
		fmt.Println("Error Happaned in Execurting the  user create query ")
		return 0, err
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
	if err != nil {
		fmt.Println("Internal Server Error while fetching data")

		return 0, err
	}

	return id, err
}

func GetUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)

	if err != nil {
		fmt.Println("Failed to Query user data from Database")

		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Created_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) ValidateCredentials() error {

	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrivedPwd string
	err := row.Scan(&u.ID, &retrivedPwd)
	if err != nil {
		fmt.Println("Invalid UserName / Password")
		return err
	}

	isVlaidPasswprd := utils.ValidatePassword(u.Password, retrivedPwd)

	if !isVlaidPasswprd {
		fmt.Println("Invalid Password (False)")

		return errors.New("invalid password")
	}
	return nil

}
