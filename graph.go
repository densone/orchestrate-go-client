package client

import (
	"encoding/json"
	"log"
	"net/url"
)

type GraphResults struct {
	Count   uint64        `json:"count"`
	Results []GraphResult `json:"results"`
}

type GraphResult struct {
	Collection string                 `json:"collection"`
	Key        string                 `json:"key"`
	Ref        string                 `json:"ref"`
	Value      map[string]interface{} `json:"value"`
}

func (client Client) GetRelations(collection string, key string, hops []string) *GraphResults {
	queryVariables := url.Values{
		"hop": hops,
	}

	req := client.newRequest("GET", collection+"/"+key+"/relations?"+queryVariables.Encode(), nil)

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	result := new(GraphResults)
	err = decoder.Decode(result)

	if err != nil {
		println(err.Error())
	}

	return result
}
