package app

import (
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func App() {
	mapurls()
}