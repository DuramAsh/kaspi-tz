package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 1000,
			IdleConnTimeout:     90 * time.Second,
		},
	}
)

func sendPerson(url string, p Person) (err error) {
	body, err := json.Marshal(p)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %v", err)
		return
	}
	defer res.Body.Close()

	io.Copy(io.Discard, res.Body)

	if res.StatusCode >= 400 {
		err = fmt.Errorf("server error [%d]: %s", res.StatusCode, p.IIN)
		log.Print(err.Error())

		return
	}

	return
}
