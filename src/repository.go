package main

import (
	"errors"
	"github.com/satori/go.uuid"
)

type UserRepository interface {
	getLastId() string
	list() []User
	get(id string) (User, error)
	add(user User) error
	update(user User) error
	remove(id string) error
}

type UserRepositoryMemory struct {
	users []User
}

var userRepo = UserRepositoryMemory{[]User{
	{"3aaa392e-c7ca-11e7-b96c-0242c0a85002", "Damiano", "Petrungaro", Address{"Berlin", "Germany"}},
	{"3aaa3994-c7ca-11e7-b96c-0242c0a85002", "Lorenzo", "D'Ianni", Address{"Milan", "Italy"}},
}}

func (repo UserRepositoryMemory) getLastId() string {
	return uuid.NewV1().String()
}

func (repo UserRepositoryMemory) list() []User {
	return repo.users
}

func (repo UserRepositoryMemory) get(id string) (User, error) {
	for _, user := range repo.users {
		if user.Id == id {
			return user, nil
		}
	}

	return User{}, errors.New("user_not_found")
}

func (repo UserRepositoryMemory) add(user User) error {
	repo.users = append(repo.users, user)

	return nil
}

func (repo UserRepositoryMemory) update(user User) error {

	for key, storedUser := range repo.users {
		if storedUser.Id == user.Id {
			repo.users[key] = user
			return nil
		}
	}

	return errors.New("user_not_found")
}

func (repo UserRepositoryMemory) remove(id string) error {
	for index, user := range repo.users {
		if user.Id == id {
			repo.users = append(repo.users[:index], repo.users[index+1:]...)
			return nil
		}
	}

	return errors.New("user_not_found")
}
