package client

import (
	"encoding/json"
	"log"
	"net/url"
)

type SearchResults struct {
	Count    uint64         `json:"count"`
	Results  []SearchResult `json:"results"`
	MaxScore float64        `json:"max_score"`
}

type SearchResult struct {
	Collection string                 `json:"collection"`
	Key        string                 `json:"key"`
	Ref        string                 `json:"ref"`
	Score      float64                `json:"score"`
	Value      map[string]interface{} `json:"value"`
}

func (client Client) Search(collection string, query string) *SearchResults {
	queryVariables := url.Values{
		"query": []string{query},
	}

	resp, err := client.doRequest("GET", collection+"?"+queryVariables.Encode(), nil)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	result := new(SearchResults)
	err = decoder.Decode(result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
