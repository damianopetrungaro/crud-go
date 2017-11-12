package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ListUsers(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(UserRepository.list(userRepo))
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	user, err := UserRepository.get(userRepo, id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = UserRepository.getLastId(userRepo)
	err := UserRepository.add(userRepo, user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	user, err := UserRepository.get(userRepo, id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	user.Id = id
	err = UserRepository.update(userRepo, user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := UserRepository.get(userRepo, id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = UserRepository.remove(userRepo, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode("Ok!")
}
