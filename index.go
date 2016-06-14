package elastisearchgo

import (
	"bytes"
	"fmt"
	"net/http"
)

type IndexSettings struct {
}

func (es *ElastiSearchClient) CreateIndex(indexname string, settings []byte) (res *http.Response, err error) {

	indexurl := es.EndPoint + "/" + indexname
	req, err := http.NewRequest("POST", indexurl, bytes.NewBuffer(settings))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create Index Status:", resp.Status)

	return resp, err
}

func (es *ElastiSearchClient) LoadDocument(indexname string, indextype string, data []byte) (res *http.Response, err error) {

	indexurl := es.EndPoint + "/" + indexname + "/" + indextype + "/" 
	req, err := http.NewRequest("POST", indexurl, bytes.NewBuffer(data))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Load Document Status:", resp.Status)

	return resp, err
}
