package elastisearchgo

import (
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
	return es
}
