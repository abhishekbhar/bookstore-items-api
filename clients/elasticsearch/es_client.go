package elasticsearch


import (
	"fmt"
	"time"
	"context"
	"github.com/olivere/elastic"
	"github.com/abhishekbhar/bookstore-items-api/logger"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface{
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://172.20.0.3:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		//elastic.SetHeaders(http.Header{
		//  "X-Caller-Id": []string{"..."},
		//}),
	)
	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}


func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}


func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
    result, err := c.client.Index().
			Index(index).
			BodyJson(doc).
			Do(ctx)
	
	if err!=nil {
		logger.Error(fmt.Sprintf("Error when tring to index document in es index %s",index), err)
		return nil, err
	}

	return result, nil
}