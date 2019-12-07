package main

import (
	"encoding/json"
	"fmt"
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

func getStanding(competitionID int) (s *standing, err error) {
	b := request(fmt.Sprintf("https://api.football-data.org/v2/competitions/%d/standings?standingType=TOTAL", competitionID))

	err = json.Unmarshal(b, &s)

	return
}
