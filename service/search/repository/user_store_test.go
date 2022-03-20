package repository

import (
	"encoding/json"
	"goblog.com/pkg/elasticsearch"
	"log"
	"testing"
	"time"
)

var es *elasticsearch.Elasticsearch
var store *UserStore

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

	store, err = NewUserStore("user", es.Client)
	if err != nil {
		log.Fatalln(err)
	}

}

func TestUserStore_CreateIndex(t *testing.T) {
	if err := store.IndicesCreate(""); err != nil {
		log.Fatalln(err)
	}
}

func TestUserStore_DeleteIndex(t *testing.T) {
	if err := store.IndicesDelete(); err != nil {
		log.Fatalln(err)
	}
}

func TestUserStore_Create(t *testing.T) {
	u := User{
		ID:       2,
		Username: "li si",
		Phone:    "13000000000",
		Email:    "lisi@email.com",
		UserInfo: UserInfo{
			UserId: 2,
			Birthday: time.Now(),
		},
	}
	if err:= store.Create(&u); err != nil {
		log.Fatalln(err)
	}
}

func TestUserStore_Update(t *testing.T) {
	var id = "1"
	u := User{
		Username: "w si",
	}

	if err:= store.Update(id, &u); err != nil {
		log.Fatalln(err)
	}
}

func TestUserStore_Index(t *testing.T) {
	u := User{
		ID:       1,
		Username: "li si",
		Phone:    "13000000000",
		Email:    "lisi@email.com",
		UserInfo: UserInfo{
			UserId: 1,
			Birthday: time.Now(),
		},
	}
	if err:= store.Index(&u); err != nil {
		log.Fatalln(err)
	}
}

func TestUserStore_Search(t *testing.T) {
	query := buildSearchQuery(searchTermUsername, "\"w si\"", 0, 2)

	res, err:= store.Search(query)
	if err != nil {
		log.Fatalln(err)
	}
	jsonB, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(jsonB))
	for _, v := range res.Hits {
		log.Println(v)
	}
}

func TestUserStore_IndicesUpdateAliases(t *testing.T) {
	res, err := store.IndicesUpdateAliases("test_alias_v1", "test_alias_v2", "test_alias")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}