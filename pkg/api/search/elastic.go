package search

import (
	"encoding/json"
	"fmt"
	elasticsearchPkg "github.com/elastic/go-elasticsearch/v8"
	"github.com/miiy/go-blog/pkg/elasticsearch"
	"log"
)

var es *elasticsearch.Elasticsearch

func init()  {
	o := &elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"},
		DiscoverNodesOnStart: true,
	}
	var err error
	es, err = elasticsearch.NewElasticsearch(o)
	if err!= nil {
		log.Fatalln(err)
	}
}

func main()  {
	Info()
	ping()
}

func Info()  {
	res, err := es.Client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	var r map[string]interface{}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print server info
	fmt.Println(r)
	for k, v := range r {
		log.Printf("%s: %v", k, v)
	}
	// Print client numbers.
	log.Printf("client: %s", elasticsearchPkg.Version)
}

func ping()  {
	res, err := es.Client.Ping()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println(res)

}