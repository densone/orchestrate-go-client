package client

import (
	"bytes"
	"io"
	"log"
)

func (client Client) GetEvents(collection string, key string, kind string) (*bytes.Buffer, error) {
	resp, err := client.doRequest("GET", collection+"/"+key+"/events/"+kind, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, newError(resp)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)

	return buf, err
}

func (client Client) PutEvent(collection string, key string, kind string, value io.Reader) error {
	resp, err := client.doRequest("PUT", collection+"/"+key+"/events/"+kind, value)

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
