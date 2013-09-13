package client

import (
	"encoding/json"
	"fmt"
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

func (client Client) GetRelations(collection string, key string, hops []string) (*GraphResults, error) {
	queryVariables := url.Values{
		"hop": hops,
	}

	resp, err := client.doRequest("GET", collection+"/"+key+"/relations?"+queryVariables.Encode(), nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, newError(resp)
	}

	decoder := json.NewDecoder(resp.Body)
	result := new(GraphResults)
	err = decoder.Decode(result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func (client Client) PutRelation(sourceCollection string, sourceKey string, kind string, sinkCollection string, sinkKey string) error {
	resp, err := client.doRequest("PUT", fmt.Sprintf("%v/%v/relations/%v/%v/%v", sourceCollection, sourceKey, kind, sinkCollection, sinkKey), nil)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		err = newError(resp)
	}

	return err
}
