package mattbas

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// Base url for api
	API = "https://insult.mattbas.org/api/insult.txt"
)

func GetInsult() string {
	resp, err := http.Get(API)

	if err != nil {
		log.Println("Error getting product ticker: ", err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error reading response body: ", err)
		return ""
	}

	return string(body)
}
