package complimentr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// Base url for api
	API = "https://complimentr.com/api"
)

func GetCompliment() string {
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

	data := make(map[string]string)
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Println("Error decoding body: ", err)
		return ""
	}

	return data["compliment"]
}
