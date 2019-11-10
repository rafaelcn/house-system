package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Serve initializes the web server
func Serve(address string, port string) {

	log.Printf("[+] Serving on %s:%s\n", address, port)

	r := mux.NewRouter()

	r.HandleFunc("/", HandleIndexPage)

	r.HandleFunc("/help", HandleHelpPage)
	r.HandleFunc("/about", HandleAboutPage)
	r.HandleFunc("/register", HandleRegisterPage)
	r.HandleFunc("/", HandleIndexPage)

	err := http.ListenAndServe(address+":"+port, nil)

	if err != nil {
		log.Fatalf("[-] Server initialization error")
	}
}

func HandleIndexPage(w http.ResponseWriter, r *http.Request) {

}

func HandleRegisterPage(w http.ResponseWriter, r *http.Request) {

}

func HandleHelpPage(w http.ResponseWriter, r *http.Request) {

}

func HandleAboutPage(w http.ResponseWriter, r *http.Request) {

}
