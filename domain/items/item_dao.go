package items

import (
	"encoding/json"
	"github.com/abhishekbhar/bookstore-items-api/clients/elasticsearch"
	"github.com/abhishekbhar/bookstore-utils-go/rest_errors"
	"github.com/abhishekbhar/bookstore-items-api/domain/queries"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item)Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("Error when trying to save item")
	}

	i.Id = result.Id
	return nil
}

func (i *Item) Get() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Get(indexItems,typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("Error when trying to get item")
	}

	if !result.Found {
		return rest_errors.NewNotFoundError("item not found")
	}

	bytes, err := result.Source.UnmarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response")
		
	}

	if err:= json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response")
		
	}
	i.Id = result.Id
	return nil
}


func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.respErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())

	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents")
	}

	return result, nil

}