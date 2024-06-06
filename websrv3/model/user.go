package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// User struct definition
type User struct {
	Id   uuid.UUID `json:"id"   validate:"required,uuid"`
	Name string    `json:"name" validate:"required"`
	Age  int       `json:"age"  validate:"gte=18,lte=130"`
}

// UserList struct definition and methods for CRUD operations
type UsersList struct {
	Users []User
}

var validate *validator.Validate

func (u *UsersList) ListAll(offset uint) ([]User, error) {
	db, err := sql.Open("sqlite3", "./db/websrv3.db")
	if err != nil {
		return []User{}, err
	}

	// pagination limit 5, offset defined by user
	const LIMIT uint = 5
	const query string = `SELECT id, name, age FROM users limit ? offset ?;`
	rows, err := db.Query(query, LIMIT, offset)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	users := u.Users
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UsersList) Create(name string, age int) (User, error) {
	user := User{
		Id:   uuid.New(),
		Name: name,
		Age:  age,
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		log.Println(err)
		return User{}, err
	}

	db, err := sql.Open("sqlite3", "./db/websrv3.db")
	if err != nil {
		return User{}, errors.New("error to open database connection")
	}

	const query string = ` INSERT INTO users (id, name, age) VALUES (?, ?, ?)`
	if _, err := db.Exec(query, user.Id, user.Name, user.Age); err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UsersList) Read(id string) (User, error) {
	db, err := sql.Open("sqlite3", "./db/websrv3.db")
	if err != nil {
		return User{}, err
	}

	_, err = uuid.Parse(id)
	if err != nil {
		return User{}, err
	}

	var user User
	var query string = `SELECT id, name, age FROM users WHERE id = ?;`
	row := db.QueryRow(query, id)
	if err = row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UsersList) Update(id string, name string, age int) (User, error) {
	db, err := sql.Open("sqlite3", "./db/websrv3.db")
	if err != nil {
		return User{}, err
	}

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return User{}, err
	}

	user := User{
		Id:   parsedUUID,
		Name: name,
		Age:  age,
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			return User{}, err
		}
	}

	var query string = `UPDATE users SET name=?, age=? WHERE id=?;`
	if _, err := db.Exec(query, user.Name, user.Age, user.Id); err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UsersList) Delete(id string) (string, error) {
	db, err := sql.Open("sqlite3", "./db/websrv3.db")
	if err != nil {
		return "", err
	}

	_, err = uuid.Parse(id)
	if err != nil {
		return "", err
	}

	var query string = `DELETE FROM users WHERE id=?;`
	if _, err := db.Exec(query, id); err != nil {
		return "", err
	}

	msg := fmt.Sprintf("deleted user with id: %s", id)
	return msg, nil
}
