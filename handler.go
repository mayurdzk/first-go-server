package main

import (
	"log"
	"net/http"
	"os"
)

type LoggingHandler struct {
	handler *http.ServeMux
}

func (mux LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: Address: %s Method: %s URL: %s  HOST: %s, UserAgent: %s\n", r.RemoteAddr, r.Method, r.URL, r.Host, r.UserAgent())
	lrw := LogResponseWriter{rw: w}
	mux.handler.ServeHTTP(lrw, r)
}

func FirstGoServerRouter(db *DB) LoggingHandler {
	configureLogEnvironment("log.txt")

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/people", http.StatusFound)
	})
	router.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		getAllPeople(w, db)
	})
	router.HandleFunc("/add-person", func(w http.ResponseWriter, r *http.Request) {
		addPerson(w, r, db)
	})

	return LoggingHandler{handler: router}
}

// Taken from: https://gist.github.com/hoitomt/c0663af8c9443f2a8294
func configureLogEnvironment(path string) {
	if path != "" {
		lf, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
		assertNotError(err)
		log.SetOutput(lf)
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // TODO: How does pipe-delimitation this work?
	} else {
		panic("Path to log file was an empty string.")
	}
}
