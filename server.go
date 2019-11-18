package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

// Serve initializes the web server
func Serve(address string, port string) {

	log.Printf("[+] Serving on %s:%s\n", address, port)

	router := mux.NewRouter()

	router.HandleFunc("/", HandleIndexPage)

	router.HandleFunc("/help", HandleHelpPage)
	router.HandleFunc("/about", HandleAboutPage)
	router.HandleFunc("/sign-up", HandleRegisterPage)
	router.HandleFunc("/login", HandleLoginPage)

	// Internal pages
	router.HandleFunc("/homepage", HandleHomePage)
	router.HandleFunc("/settings", HandleSettingsPage)
	router.HandleFunc("/user-settings", HandleUserSettingsPage)
	router.HandleFunc("/acessory", HandleAcessoryPage)

	// REST API
	router.HandleFunc("/v1/login", HandleLogin)
	router.HandleFunc("/v1/logout", HandleLogout)

	router.HandleFunc("/v1/user/{id}", HandleUser)
	router.HandleFunc("/v1/user/{action}", HandleUser)

	router.HandleFunc("/v1/people", HandlePeople)

	router.HandleFunc("/v1/object/{id}", HandleObject)
	router.HandleFunc("/v1/object/{action}", HandleObject)
	router.HandleFunc("/v1/object/{[0-9]}/{action}", HandleObject)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",
		http.FileServer(http.Dir("./html/assets/"))))
	router.Use(MiddlewareStaticFiles)

	err := http.ListenAndServe(address+":"+port, router)

	if err != nil {
		log.Fatalf("[!] Server initialization error. Reason %v", err)
	}
}

// MiddlewareStaticFiles fgn
func MiddlewareStaticFiles(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only do work if the request is for assets
		if strings.Contains(r.URL.String(), "assets") {
			if path.Ext(r.URL.String()) == "" || r.Header.Get("Referer") == "" {
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}

// HandleIndexPage ...
func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/external.html",
		"html/pages/external/index.html",
	))

	d := PageData {
		Title: "Início",
	}

	t.Execute(w, d)
}

// HandleLoginPage ...
func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/external.html",
		"html/pages/external/login.html",
	))

	d := PageData {
		Title: "Login",
	}

	t.Execute(w, d)
}

// HandleRegisterPage ...
func HandleRegisterPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/external.html",
		"html/pages/external/register.html",
	))

	d := PageData {
		Title: "Registrar-se",
	}

	t.Execute(w, d)
}

// HandleAboutPage ...
func HandleAboutPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/external.html",
		"html/pages/external/about.html",
	))

	d := PageData {
		Title: "Sobre",
	}

	t.Execute(w, d)
}

// Internal pages are defined below

// HandleHomePage ...
func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/internal.html",
		"html/pages/internal/home.html",
	))

	t.Execute(w, nil)
}

// HandleSettingsPage ...
func HandleSettingsPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/internal.html",
		"html/pages/internal/settings.html",
	))

	t.Execute(w, nil)
}

// HandleSettingsPage ...
func HandleUserSettingsPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/internal.html",
		"html/pages/internal/user_settings.html",
	))

	t.Execute(w, nil)
}

// HandleAddAcessoryPage ...
func HandleAddAcessoryPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/internal.html",
		"html/pages/internal/add_acessory.html",
	))

	t.Execute(w, nil)
}

// HandleAcessoryPage ...
func HandleAcessoryPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/internal.html",
		"html/pages/internal/acessory.html",
	))

	t.Execute(w, nil)
}

// HandleHelpPage ...
func HandleHelpPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/internal.html",
		"html/pages/internal/help.html",
	))

	t.Execute(w, nil)
}
