package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dragorific/makeuoft_wildfirepredictor/setup"
	"github.com/gorilla/mux"
)

func main() {
	//gets state file
	s := setup.GetMainState("api engine")

	//creates new router for api
	router := mux.NewRouter()
	router.StrictSlash(true)

	//sets up api subrouter
	api := router.PathPrefix("/api/").Subrouter()

	//Lets user know if route is working
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
	})

	setUpRoutes(s, router, api)

	server := &http.Server{
		Addr:         ":6060",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start HTTP server
	s.Log.Info("http endpoint now active on :6060")
	err := server.ListenAndServe()
	if err != nil {
		s.Log.Fatal(err)
	}
}

func setUpRoutes(s *setup.State, router *mux.Router, api *mux.Router) {

	api.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Success")
	})
}
