package services

import (
	"net/http"
	"github.com/abhishekbhar/bookstore-items-api/domain/items"
	"github.com/abhishekbhar/bookstore-utils-go/rest_errors"
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
	return nil, rest_errors.NewRestError("implement me!!", http.StatusNotImplemented, "not_implemented")
}

func (s *itemService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!!", http.StatusNotImplemented, "not_implemented")
}



