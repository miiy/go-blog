package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type UserStore struct {
	indexName string
	client    *elasticsearch.Client
}

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)

type User struct {
	ID int64	      `json:"id"`
	Username string   `json:"username"`
	Phone string      `json:"phone"`
	Email string      `json:"email"`
	UserInfo UserInfo `json:"user_info"`
}

type UserInfo struct {
	UserId int64       `json:"user_id"`
	Bio string         `json:"bio"`
	Age int64          `json:"age"`
	Birthday time.Time `json:"birthday"`
	Interests []string `json:"interests"`
}

func NewUserStore(indexName string, client *elasticsearch.Client) (*UserStore, error) {
	s := &UserStore{
		indexName,
		client,
	}
	return s, nil
}

const indexMapping = `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "long"
      },
      "username": {
        "type": "text",
        "analyzer": "ik_max_word",
        "fields": {
          "raw": {
            "type": "keyword"
          }
        }
      },
      "passowrd": {
        "type": "keyword",
        "index": false
      },
      "email": {
        "type": "text",
        "analyzer": "ik_max_word",
        "fields": {
          "raw": {
            "type": "keyword"
          }
        }
      },
      "phone": {
        "type": "keyword"
      },
      "status": {
        "type": "short"
      }
    }
  }
}
`

// IndicesCreate creates an index with optional settings and mappings.
//
func (s *UserStore) IndicesCreate(mapping string) error {
	res, err := s.client.Indices.Create(s.indexName, s.client.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}
	return nil
}

// IndicesDelete deletes an index.
//
func (s *UserStore) IndicesDelete() error {
	res, err := s.client.Indices.Delete([]string{s.indexName})
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}
	return nil
}

// IndicesPutAlias creates or updates an alias.
//
func (s *UserStore) IndicesPutAlias(name string) error {
	res, err := s.client.Indices.PutAlias([]string{s.indexName}, name)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}
	return nil
}


func (s *UserStore) IndicesUpdateAliases(oldIndex, newIndex, alias string) (bool, error) {
	var q = `
{
  "actions": [
	{"remove": {"index": "%s", "alias": "%s"}},
    {"add": {"index": "%s", "alias": "%s"}}
  ]
}
`
	q = fmt.Sprintf(q, oldIndex, alias, newIndex, alias )
	res, err := s.client.Indices.UpdateAliases(strings.NewReader(q))
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return false, fmt.Errorf("error: %s", res)
	}

	return true, nil
}

func (s *UserStore) Bulk(users []*User) {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		NumWorkers:          runtime.NumCPU(),
		FlushBytes:          0,
		FlushInterval:       0,
		Client:              s.client,
		Index:               s.indexName,
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	var countSuccessful uint64
	for _, item := range users {
		data, err := json.Marshal(item)
		if err != nil {
			log.Fatalf("Cannot encode article %d: %s", item.ID, err)
		}

		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",
				DocumentID: strconv.FormatInt(item.ID, 10),
				Body: bytes.NewReader(data),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}

	}
	// close the indexer
	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	biStats := bi.Stats()
	log.Println(biStats)

}

// Create creates a new document in the index.
//
// Returns a 409 response when a document with a same ID already exists in the index.
//
func (s *UserStore) Create(item *User) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(item); err != nil {
		return err
	}

	res, err := s.client.Create(s.indexName, strconv.FormatInt(item.ID, 10), &buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 409 {
		return ErrConflict
	}
	if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}

	return nil
}

// Index creates or updates a document in an index.
//
func (s *UserStore) Index(item *User) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(item); err != nil {
		return err
	}

	itemId := strconv.FormatInt(item.ID, 10)
	res, err := s.client.Index(s.indexName, &buf, s.client.Index.WithDocumentID(itemId))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}

	return nil
}

// Exists returns information about whether a document exists in an index.
func (s *UserStore) Exists(id string) (bool, error) {
	res, err := s.client.Exists(s.indexName, id)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		return true, nil
	case 404:
		return false, nil
	default:
		return false, fmt.Errorf("error: %s", res)
	}

}

// Get returns a document.
//
func (s *UserStore) Get(id string) (*User, error) {
	res, err := s.client.Get(s.indexName, id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, ErrNotFound
	}
	if res.IsError() {
		return nil, fmt.Errorf("error: %s", res)
	}

	var user User
	if err = json.NewDecoder(res.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates a document with a script or partial document.
//
func (s *UserStore) Update(id string, i *User) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(i); err != nil {
		return err
	}

	body := bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, &buf)))

	res, err := s.client.Update(s.indexName, id, body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		return nil
	case 404:
		return ErrNotFound
	default:
		return fmt.Errorf("error: %s", res)
	}
}

// Delete removes a document from the index.
//
func (s *UserStore) Delete(ctx context.Context, id string) error {
	res, err := s.client.Delete(s.indexName, id, s.client.Delete.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		return nil
	case 404:
		return ErrNotFound
	default:
		return fmt.Errorf("error: %s", res)
	}
}



const searchAll = `
{
    "query": {
        "match_all": {}
    },
	"from": {from},
    "size": {size},
	"sort": {
		"id": "desc",
		"_doc" : "asc"
	}
}
`

const searchTermUsername = `
{
    "query": {
        "term": {
			"username": {
				"value": "{username}"
			}
		}
    }
}
`

type SearchResults struct {
	Total int   `json:"total"`
	Hits []*Hit `json:"hits"`
}

type Hit struct {
	ID         string          `json:"id"`
	Source     User            `json:"source"`
	Sort       []interface{}   `json:"sort"`
	Highlights json.RawMessage `json:"highlight"`
}

func buildSearchQuery(query, username string, from, size int64) string {
	if from < 0 {
		from = 0
	}
	if size < 0 {
		size = 20
	}
	query = strings.Replace(query, "{from}", strconv.FormatInt(from, 10), 1)
	query = strings.Replace(query, "{size}", strconv.FormatInt(size, 10), 1)
	query = strings.Replace(query, "{username}", username, 1)
	return query
}

func (s *UserStore) Search(query string) (*SearchResults, error) {
	var results SearchResults

	body := strings.NewReader(query)

	res, err := s.client.Search(
		s.client.Search.WithIndex(s.indexName),
		s.client.Search.WithBody(body),
		//s.client.Search.WithFrom(from),
		//s.client.Search.WithSize(size),
		//s.client.Search.WithSort(sort),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error: %s", res)
	}

	type envelopeResponse struct {
		Took int
		Hits struct {
			Total struct {
				Value int
			}
			Hits []struct {
				ID         string          `json:"_id"`
				Source     json.RawMessage `json:"_source"`
				Highlights json.RawMessage `json:"highlight"`
				Sort       []interface{}   `json:"sort"`
			}
		}
	}

	var r envelopeResponse
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return &results, err
	}

	results.Total = r.Hits.Total.Value

	for _, hit := range r.Hits.Hits {
		var h Hit

		h.ID = hit.ID
		if err = json.Unmarshal(hit.Source, &h.Source); err != nil {
			return &results, err
		}
		h.Sort = hit.Sort
		if len(hit.Highlights) > 0 {
			if err = json.Unmarshal(hit.Highlights, &h.Highlights); err != nil {
				return &results, err
			}
		}

		results.Hits = append(results.Hits, &h)
	}

	return &results, nil
}
