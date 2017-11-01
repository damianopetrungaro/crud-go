package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(list())
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	user, error := get(id)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := User{Id: getLastId()}
	json.NewDecoder(r.Body).Decode(&user)
	error := add(user)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	user, error := get(id)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	user.Id = id
	error = update(user)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, error := get(id)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error)
		return
	}

	error = remove(id)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode("Ok!")
}
