package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HandleUser
func HandleUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)

	db := PGSQLConnect()

	if vars["id"] != "" { 

	} else if vars["action"] != "" {
		switch vars["action"] {
		case "add":

		case "delete":

		default:
		}
	}
}

// HandlePeople is a handle that accepts requests to manage a group of people
func HandlePeople(w http.ResponseWriter, r *http.Request) {
	
}

// HandleHardware
func HandleObject(w http.ResponseWriter, r *http.Request) {

}
