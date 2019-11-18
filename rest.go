package main

import (
	"database/sql"
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
		row := db.Query(SelectUser, []interface{}{vars["id"]})
		var user User

		for row.Next() {
			err := row.Scan(&user.ID, &user.Name, &user.Mail, &user.Phone,
				&user.Birth,
				&user.Type)

			switch err {
			case sql.ErrNoRows:
				response.Status = StatusError

				errorMessage := fmt.Sprintf("Database request error, "+
					"notify the developer about %v.", err.Error())

				e := Error{
					Code:        ErrorDatabaseResponse,
					Description: errorMessage,
				}

				response.Content = e
			case nil:
				response.Status = StatusOk
				response.Content = user
			}
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

	rows := db.Query(FetchUsers, []interface{}{})
	
	var response Response
	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Mail, &user.Phone,
			&user.Birth,
			&user.Type)

		switch err {
		case sql.ErrNoRows:
			response.Status = StatusError

			errorMessage := fmt.Sprintf("Database request error, "+
				"notify the developer about %v.", err.Error())

			e := Error{
				Code:        ErrorDatabaseResponse,
				Description: errorMessage,
			}

			response.Content = e
			break
		case nil:
			response.Status = StatusOk

			users = append(users, user)

			response.Content = users
		}
	}
	message, err := json.Marshal(response)

	if err != nil {
		Report500(&w,
			fmt.Sprintf("[!] Error encoding data to json. Reason %v", err))
	}

	w.Write(message)
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
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	db := PGSQLConnect()

	db.QueryRow(Login, []interface{}{username, password})

}

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
