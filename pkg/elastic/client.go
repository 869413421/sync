package elastic

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
	EsClient, err = elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(esConfig.Nodes...), elastic.SetBasicAuth(esConfig.Name, esConfig.Password))
	if err != nil {
		logger.Danger("connectEs error:", err)
		panic(err)
	}
}

func Create() {
	info, code, err := EsClient.Ping("http://47.94.155.227:9200").Do(context.Background())
	fmt.Println(info)
	fmt.Println(code)
	fmt.Println(err)
}
