package client

import (
	"bytes"
	"io"
	"log"
)

func (client Client) Get(collection string, key string) (*bytes.Buffer, error) {
	resp, err := client.doRequest("GET", collection+"/"+key, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, newError(resp)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf, nil
}

func (client Client) Put(collection string, key string, value io.Reader) error {
	resp, err := client.doRequest("PUT", collection+"/"+key, value)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		err = newError(resp)
	}

	return err
}
