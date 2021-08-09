package elasticsearch_client

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"sync/config"
	"sync/pkg/logger"
)

var EsClient *elastic.Client
var esConfig config.ElasticSearch

func init() {
	//1.初始化ES
	esConfig := config.LoadConfig().ElasticSearch
	fmt.Println(esConfig)
	var err error
	EsClient, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esConfig.Nodes...), elastic.SetBasicAuth(esConfig.Name, esConfig.Password))
	if err != nil {
		logger.Danger("connectEs error:", err)
		panic(err)
	}
}

// Bulk 批量请求ES
func Bulk(reqs []elastic.BulkableRequest) (*elastic.BulkResponse, error) {
	if len(reqs) <= 0 {
		return nil, nil
	}
	bulk := EsClient.Bulk()
	bulk.Add(reqs...)
	do, err := bulk.Do(context.Background())
	fmt.Println(do)
	fmt.Println(err)
	return do, err
}
