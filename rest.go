package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// HandleUser
func HandleUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)

	db := PGSQLConnect()

	var response Response

	if vars["action"] != "" {
		switch vars["action"] {
		case "fetch":
			if vars["id"] == "" {
				response = IncompleteRequest()
			} else {
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
			}
		case "sign-up":
			name := r.Form.Get("name")
			email := r.Form.Get("email")
			password := r.Form.Get("password")
			phone := r.Form.Get("phone")
			birth := r.Form.Get("birth")

			if name == "" || email == "" || password == "" || phone == "" || birth == "" {
				response = IncompleteRequest()
			} else {
				// Format phone
				phone = strings.Replace(phone, "(", "", 1)
				phone = strings.Replace(phone, ")", "", 1)
				phone = strings.Replace(phone, "-", "", 1)
				phone = strings.Replace(phone, " ", "", 1)

				// Format birth date
				date := strings.Split(birth, "/")
				year := date[2]
				date[2] = date[0]
				date[0] = year
				birth = strings.Join(date, "-")

				db.Execute(AddUser, []interface{}{name, email, password, phone,
					birth})

				response.Status = StatusOk
				response.Content = "Você foi registrado com sucesso."
			}
		case "update":
			id := r.Form.Get("id")
			name := r.Form.Get("name")
			email := r.Form.Get("email")
			password := r.Form.Get("password")
			phone := r.Form.Get("phone")
			birth := r.Form.Get("birth")

			db.Execute(UpdateUser, []interface{}{id, name, email, password,
				phone, birth})
		case "delete":
			id := r.Form.Get("id")
			db.Execute(RemoveUser, []interface{}{id})
		default:
		}
	}

	respond(&w, response)
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

	respond(&w, response)
}

// HandleObject ...
func HandleObject(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)

	db := PGSQLConnect()

	var response Response
	var err error

	if vars["action"] != "" {
		switch vars["action"] {
		case "fetch":
			if vars["id"] == "" {
				response = IncompleteRequest()
			} else {
				id := vars["id"]

				row := db.Query(SelectObject, []interface{}{id})
				var object Object

				for row.Next() {
					err = row.Scan(&object.ID, &object.Name, &object.Status, &object.Type,
						&object.House, &object.Intensity, &object.Volume, &object.Distance,
						&object.Temperature)

					if err != nil {
						errorMessage := fmt.Sprintf("Database request error, "+
							"notify the developer about %v.", err.Error())

						e := Error{
							Code:        ErrorDatabaseResponse,
							Description: errorMessage,
						}

						response.Status = StatusError
						response.Content = e
					} else {
						response.Status = StatusOk
						response.Content = object
					}
				}
			}
		case "update":
			objectID := r.Form.Get("code")
			objectName := r.Form.Get("name")
			objectStatus := r.Form.Get("status")
			objectType := r.Form.Get("type")
			objectNType := -1

			switch objectType {
			case "light":
				objectNType = 1
			case "sound":
				objectNType = 2
			case "sensor":
				objectNType = 3
			case "air-conditioner":
				objectNType = 4
			}

			_ = db.Execute(UpdateObject, []interface{}{objectName, objectStatus,
				objectNType, objectID})

			if objectID == "" {
				response = IncompleteRequest()
			} else {

			}
		case "delete":
			id := r.Form.Get("code")

			if id == "" {
				response = IncompleteRequest()
			} else {
				_ = db.Execute(RemoveObject, []interface{}{id})
			}
			
		case "add":
			objectID := r.Form.Get("code")
			objectName := r.Form.Get("name")
			objectType := r.Form.Get("type")
			objectNType := -1

			switch objectType {
			case "light":
				objectNType = 1
			case "sound":
				objectNType = 2
			case "sensor":
				objectNType = 3
			case "air-conditioner":
				objectNType = 4
			}

			_ = db.Execute(AddObject, []interface{}{objectID,
				objectName, false, objectNType, 1, 0.0, 0.0, 0.0, 0.0})

			response.Status = StatusOk
		}

		respond(&w, response)
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

			e := Error{
				Code:        ErrorDatabaseResponse,
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

	respond(&w, response)
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
	var e Error = Error{Code: ErrorNotAuthorized}

	err := row.Scan(&user.ID, &user.Name, &user.Mail, &user.Username,
		&user.Password, &user.Phone, &user.Birth, &user.Type)

	switch err {
	case sql.ErrNoRows:
		e.Description = "Email or password incorrect"
		response.Content = e
		
	case nil:
		// TODO: Create an user session
		response.Content = user.ID
	default:
		log.Printf("[!] Unknown error: %v", err)
	}

	respond(&w, response)
}

// HandleInvite ...
func HandleInvite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)

	db := PGSQLConnect()

	var response Response

	switch vars["action"] {
	case "new":
		email := r.Form.Get("email")
		db.Execute(InviteNew, []interface{}{email})

		response.Status = StatusOk
	case "fetch":
	}

	respond(&w, response)
}

// HandleLogout ...
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: Terminate user session

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Responds the JSON to the caller
func respond(w *http.ResponseWriter, response Response) {
	message, err := json.Marshal(response)

	if err != nil {
		Report500(w,
			fmt.Sprintf("[!] Couldn't encode data in json. Reason %v", err))
	} else {
		(*w).Write(message)
	}
}

// IncompleteRequest ...
func IncompleteRequest() Response {
	var response Response

	e := Error{
		Code:        ErrorIncompleteRequest,
		Description: "The given endpoint requires data as specified.",
	}

	response.Status = StatusError
	response.Content = e

	return response
}
