package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main()  {
	tagList()
}


func tagList()  {

	var body string

	resp, err := http.Get("http://localhost:8051/v1/tag?page=1&per_page=5")

	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n\n", resp.StatusCode, body)
}

func tagGet()  {

	var body string

	resp, err := http.Get("http://localhost:8051/v1/tag/1")

	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n\n", resp.StatusCode, body)
}

func tagCreate()  {

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)

	var body string

	resp, err := http.Post("http://localhost:8051/v1/tag", "application/json", strings.NewReader(fmt.Sprintf(`
		{
			"tag": {
				"name":"title (%s)",
				"description":"description (%s)"
			}
		}
	`, pfx, pfx)))

	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n\n", resp.StatusCode, body)
}
