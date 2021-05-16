package services

import (
	"net/http"
	"github.com/abhishekbhar/bookstore-items-api/domain/items"
	"github.com/abhishekbhar/bookstore-utils-go/rest_errors"
	"github.com/abhishekbhar/bookstore-items-api/domain/queries"
)

var (
	ItemService itemServiceInterface = &itemService{}
)


type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemService struct{}


func (s *itemService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	if err:= item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(id string) (*items.Item, *rest_errors.RestErr) {

	item := items.Item{Id: id}

	if err := item.Get() ; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Search(query queries.EsQuery) {
	dao := items.item{}
	dao.Search(query)
}


