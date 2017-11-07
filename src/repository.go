package main

import (
	"errors"
	"github.com/satori/go.uuid"
)

var users = []User{
	{getLastId(), "John", "Doe", Address{"City X", "State X"}},
	{getLastId(), "Koko", "Doe", Address{"City Z", "State Y"}},
}

func getLastId() string {
	return uuid.NewV1().String()
}

func list() []User {
	return users
}

func get(id string) (User, error) {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}

	return User{}, errors.New("user_not_found")
}

func add(user User) error {
	users = append(users, user)

	return nil
}

func update(user User) error {

	for key, stored_user := range users {
		if stored_user.Id == user.Id {
			users[key] = user
			return nil
		}
	}

	return errors.New("user_not_found")
}

func remove(id string) error {
	for index, user := range users {
		if user.Id == id {
			users = append(users[:index], users[index+1:]...)
			return nil
		}
	}

	return errors.New("user_not_found")
}
