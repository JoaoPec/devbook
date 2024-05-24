package database

import (
	"database/sql"
	"devbook/models"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func InsertUser(user models.User) {

	db, err := sql.Open("sqlite3", "devbook.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	userExists, err := db.Prepare(`SELECT id FROM users WHERE email = ?`)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := userExists.Query(user.Email)
	if err != nil {
		fmt.Println(err)
		return
	}

	if rows.Next() {
		fmt.Println("User already exists")
		return
	}

	query, err := db.Prepare(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = query.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User inserted successfully")
}

func SearchUserById(id string) (error) {

	db, err := sql.Open("sqlite3", "devbook.db")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	row := db.QueryRow(`SELECT * FROM users WHERE id = ?`, id)

	var user models.User

	err = row.Scan(&user.Name, &user.Email, &user.Password, &user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found")
			return err
		}
		fmt.Println(err)

		return err

	}

	fmt.Println("This is the user: ", user)

    return nil
}
