package client

import (
	"bytes"
	"errors"
	"io"
	"log"
)

func (client Client) GetEvents(collection string, key string, kind string) *bytes.Buffer {
	req := client.newRequest("GET", collection+"/"+key+"/events/"+kind, nil)

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf
}

func (client Client) PutEvent(collection string, key string, kind string, value io.Reader) error {
	req := client.newRequest("PUT", collection+"/"+key+"/events/"+kind, value)

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		err = errors.New(resp.Status)
	}

	return err
}
