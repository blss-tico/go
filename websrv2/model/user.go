package model

import (
	"errors"

	"github.com/google/uuid"
)

// User struct definition
type User struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

// UserList struct definition and methods for CRUD operations
type UsersList struct {
	Users []User
}

func (u *UsersList) ListAll() []User {
	return u.Users
}

func (u *UsersList) Create(name string, age int) User {
	user := User{
		Id:   uuid.New(),
		Name: name,
		Age:  age,
	}

	u.Users = append(u.Users, user)
	return user
}

func (u *UsersList) Read(id string) (User, error) {
	for _, user := range u.Users {
		if user.Id.String() == id {
			return user, nil
		}
	}

	err := errors.New("user not found to read")
	return User{}, err
}

func (u *UsersList) Update(id string, name string, age int) (User, error) {
	for i, user := range u.Users {
		if user.Id.String() == id {
			u.Users[i].Name = name
			u.Users[i].Age = age
			return u.Users[i], nil
		}
	}

	err := errors.New("user not found to update")
	return User{}, err
}

func (u *UsersList) Delete(id string) error {
	for i, user := range u.Users {
		if user.Id.String() == id {
			u.Users = append(u.Users[:i], u.Users[i+1:]...)
			return nil
		}
	}

	return errors.New("user not found to update")
}
