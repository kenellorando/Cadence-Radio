// routes.go
// Just the router

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/aria1/search", handleARIA1Search()).Methods("POST")
	r.HandleFunc("/api/aria1/request", handleARIA1Request()).Methods("POST")
	r.HandleFunc("/api/aria1/library", handleARIA1Library()).Methods("GET")
	r.HandleFunc("/api/aria2/request", handleARIA2Request()).Methods("POST")
	r.HandleFunc("/", handleServeRoot()).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public/static/")))).Methods("GET")
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css/")))).Methods("GET")
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js/")))).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(handleServe404())

	return r
}
