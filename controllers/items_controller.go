package controllers


import (
	"fmt"
	"net/http"
	"github.com/abhishekbhar/bookstore-items-api/domain/items"
	"github.com/abhishekbhar/bookstore-oauth-go/oauth"
	"github.com/abhishekbhar/bookstore-items-api/services"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return error json to the user
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemService.Create(item)	
	if err != nil {
		// TODO: return error json to the user
	}

	fmt.Println(result)
	// TODO: Return created item as json  with http status 201 - Created 
}

func Get(w http.ResponseWriter, r *http.Request) {

}