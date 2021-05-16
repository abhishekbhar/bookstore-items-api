package controllers


import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/abhishekbhar/bookstore-items-api/utils/http_utils"
	"github.com/abhishekbhar/bookstore-items-api/domain/items"
	"github.com/abhishekbhar/bookstore-oauth-go/oauth"
	"github.com/abhishekbhar/bookstore-items-api/services"
	"github.com/abhishekbhar/bookstore-utils-go/rest_errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {}

func (i *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return 
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, *respErr)
		return
	}

	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, *respErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, err := services.ItemService.Create(itemRequest)	
	if err != nil {
		http_utils.RespondError(w, *err)
		return
	}

	http_utils.RespondJson(w,http.StatusCreated, result)
}

func (i *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}