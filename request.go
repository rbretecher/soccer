package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func request(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	println("DEBUG TOKEN")
	println(os.Getenv("TOKEN"))

	req.Header.Set("X-Auth-Token", os.Getenv("TOKEN"))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
