package app



import (
	"net/http"
	"github.com/abhishekbhar/bookstore-items-api/controllers"
)


func mapurls() {
	router.HandleFunc("/items", controllers.Create).Methods(http.MethodPost)
}