package client

import (
	"bytes"
	"errors"
	"io"
	"log"
)

func (client Client) Get(collection string, key string) *bytes.Buffer {
	resp, err := client.doRequest("GET", collection+"/"+key, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf
}

func (client Client) Put(collection string, key string, value io.Reader) error {
	resp, err := client.doRequest("PUT", collection+"/"+key, value)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		err = errors.New(resp.Status)
	}

	return err
}
