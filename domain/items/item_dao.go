package items

import (
	"github.com/abhishekbhar/bookstore-items-api/clients/elasticsearch"
	"github.com/abhishekbhar/bookstore-utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item)Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("Error when trying to save item")
	}

	i.Id = result.Id
	return nil
}