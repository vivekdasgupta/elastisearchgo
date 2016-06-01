package elastisearchgo

import (
	"bytes"
	"fmt"
	"net/http"
)

type IndexSettings struct {
}

func (es *ElastiSearchClient) CreateIndex(indexname string, settings []byte) (res int, err error) {

	indexurl := es.EndPoint + "/" + indexname
	req, err := http.NewRequest("POST", indexurl, bytes.NewBuffer(settings))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create Index Status:", resp.Status)

	return
}
