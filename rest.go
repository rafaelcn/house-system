package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleUser
func HandleUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)

	db := PGSQLConnect()

	var response Response

	if vars["id"] != "" {
		row := db.QueryRow(SelectUser, []interface{}{vars["id"]})

		if row != nil {
			var user User

			row.Scan(&user.ID)
			row.Scan(&user.Name)
			row.Scan(&user.Mail)
			row.Scan(&user.Phone)
			row.Scan(&user.Birth)
			row.Scan(&user.Type)

			response.Status = StatusOk
			response.Content = user
		} else {
			response.Status = StatusError
			
			e := Error {
				Code: ErrorDatabaseResponse,
				Description: "Database request error, notify the developer.",
			}

			response.Content = e
		}
	} else if vars["action"] != "" {
		switch vars["action"] {
		case "update":

		case "delete":

		default:
		}
	}

	message, err := json.Marshal(response)

	if err != nil {
		Report500(&w,
			fmt.Sprintf("[!] Error encoding data to json. Reason %v", err))
	}

	w.Write(message)
}

// HandlePeople is a handle that accepts requests to manage a group of people
func HandlePeople(w http.ResponseWriter, r *http.Request) {
	db := PGSQLConnect()

	results := db.Query(FetchUsers, []interface{}{})

	for results.Next() {

	}
}

// HandleHardware
func HandleObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//db := PGSQLConnect()

	if vars["id"] != "" {

	} else if vars["action"] != "" {
		switch vars["action"] {
		case "add":

		case "delete":

		default:
		}
	}
}

// HandleLogin ...
func HandleLogin(w http.ResponseWriter, r *http.Request) {}

// HandleLogout ...
func HandleLogout(w http.ResponseWriter, r *http.Request) {}

// Report500 reports an Internal Server Error (HTTP_500) to the client
func Report500(w *http.ResponseWriter, message string) {
	log.Printf(message)

	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Write([]byte("500 Internal Server Error"))
}

// Report503 reports a Service Unavailable Error (HTTP_503) to the client
func Report503(w *http.ResponseWriter, message string) {
	log.Printf(message)

	(*w).WriteHeader(http.StatusServiceUnavailable)
	(*w).Write([]byte("503 Service Unavailable"))
}
