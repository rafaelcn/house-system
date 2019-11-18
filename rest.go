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
		case "sign-up":
			name := r.Form.Get("name")
			email := r.Form.Get("email")
			password := r.Form.Get("password")
			phone := r.Form.Get("phone")
			birth := r.Form.Get("birth")

			db.Execute(AddUser, []interface{}{name, email, password, phone,
				birth})

			// Make sure the user will login
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		case "update":
			name := r.Form.Get("name")
			email := r.Form.Get("email")
			password := r.Form.Get("password")
			phone := r.Form.Get("phone")
			birth := r.Form.Get("birth")

			db.Execute(AddUser, []interface{}{name, email, password, phone,
				birth})
		case "delete":
			id := r.Form.Get("id")
			db.Execute(RemoveUser, []interface{}{id})
		case "invite":

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

// HandleObjects ...
func HandleObjects(w http.ResponseWriter, r *http.Request) {
	db := PGSQLConnect()

	rows := db.Query(FetchObjects, []interface{}{})

	var response Response
	var objects []Object
	var err error

	for rows.Next() {
		var object Object

		err = rows.Scan(&object.ID, &object.Name, &object.Status, &object.Type,
			&object.House, &object.Intensity, &object.Volume, &object.Distance,
			&object.Temperature)

		if err != nil {
			errorMessage := fmt.Sprintf("Database request error, "+
			"notify the developer about %v.", err.Error())

			e := Error {
				Code: ErrorDatabaseResponse,
				Description: errorMessage,
			}

			response.Status = StatusError
			response.Content = e
		}

		objects = append(objects, object)
	}

	if err == nil {
		response.Status = StatusOk
		response.Content = objects
	}

	message, err := json.Marshal(response)

	if err != nil {
		Report500(&w,
			fmt.Sprintf("[!] Couldn't encode data in json. Reason %v", err))
	} else {
		w.Write(message)
	}
}

// HandleLogin ...
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	db := PGSQLConnect()

	row := db.db.QueryRow(Login, email, password)

	var user User
	var response Response = Response{Status: StatusOk}

	err := row.Scan(&user.ID, &user.Name, &user.Mail, &user.Password,
		&user.Phone, &user.Birth, &user.Type)

	switch err {
	case sql.ErrNoRows:
		response.Content = "Email or password incorrect"

		message, err := json.Marshal(response)

		if err != nil {
			Report500(&w,
				fmt.Sprintf("[!] Couldn't encode data to json. Reason %v",
					err))
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			w.Write(message)
		}
		break
	case nil:
		// TODO: Create an user session
		http.Redirect(w, r, "/homepage", http.StatusSeeOther)
		break
	default:
		log.Printf("[!] Unknown error: %v", err)
	}
}

// HandleLogout ...
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: Terminate user session

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

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
