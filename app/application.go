package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/abhishekbhar/bookstore-items-api/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func App() {
	elasticsearch.Init()
	mapurls()

	srv := &http.Server{
        Addr:         "127.0.0.1:8080",
        // // Good practice to set timeouts to avoid Slowloris attacks.
        // WriteTimeout: time.Second * 15,
        // ReadTimeout:  time.Second * 15,
        // IdleTimeout:  time.Second * 60,
        Handler: router, // Pass our instance of gorilla/mux in.
    }

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}