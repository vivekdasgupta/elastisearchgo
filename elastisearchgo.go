package elastisearchgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type ElastiSearchClient struct {
	EndPoint string
}

func NewElastiSearchClient(url string) (es *ElastiSearchClient, err error) {
	es := new(ElastiSearchClient)
	es.EndPoint = url
	return es
}

func (es *ElastiSearchClient) genRequest(index string, indextype string, settings map[string]string) (request string, err error) {

	params := url.Values{}
	for i, j := range settings {
		params.Set(i, j)
	}
	request = fmt.Sprintf("%s/%s/%s?%s", es.EndPoint, url.QueryEscape(index), url.QueryEscape(indextype), params.Encode())
	return request
}

func prettyPrintJson(str string) string {
	var pstr bytes.Buffer
	err := json.Indent(&pstr, []byte(str), "", "\t")
	if err != nil {
		fmt.Printf("%s", str)
	}
	fmt.Printf("%s", pstr)
}
