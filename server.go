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
	router.HandleFunc("/register", HandleRegisterPage)

	// REST API
	//r.HandleFunc("/v1/people/")
	//r.HandleFunc("/v1/people/{action}")

	//r.HandleFunc("/v1/hardware/")
	//r.HandleFunc("/v1/hardware/{action}")
	//r.HandleFunc("/v1/hardware/{[0-9]}/{action}")

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

// HandleIndexPage
func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"html/templates/external.html",
		"html/pages/index.html",
	))

	t.Execute(w, nil)
}

func HandleRegisterPage(w http.ResponseWriter, r *http.Request) {

}

func HandleHelpPage(w http.ResponseWriter, r *http.Request) {

}

func HandleAboutPage(w http.ResponseWriter, r *http.Request) {

}
