package app



import (
	"net/http"
	"github.com/abhishekbhar/bookstore-items-api/controllers"
)


func mapurls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}