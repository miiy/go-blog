package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

type Config elasticsearch.Config

type Elasticsearch struct {
	Client *elasticsearch.Client
	Options *Config
}

func NewElasticsearch(o *Config) (*Elasticsearch, error) {
	cfg := elasticsearch.Config{
		Addresses: o.Addresses,
		DiscoverNodesOnStart: o.DiscoverNodesOnStart,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return &Elasticsearch{
		Client: es,
		Options: o,
	}, nil
}
