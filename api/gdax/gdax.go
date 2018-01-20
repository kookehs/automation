package gdax

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// Base url for api
	API = "https://api.gdax.com"

	// Trading pair IDs
	BTCUSD = "BTC-USD"
)

type ProductTicker struct {
	Ask     string `json:ask`
	Bid     string `json:bid`
	Price   string `json:price`
	Size    string `json:size`
	Time    string `json:time`
	TradeID uint64 `json:trade_id`
	Volume  string `json:volume`
}

func (pt *ProductTicker) String() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("A:\t")
	buffer.WriteString(pt.Ask)
	buffer.WriteString("\nB:\t")
	buffer.WriteString(pt.Bid)
	buffer.WriteString("\nP:\t")
	buffer.WriteString(pt.Price)
	buffer.WriteString("\nS:\t")
	buffer.WriteString(pt.Size)
	buffer.WriteString("\nV:\t")
	buffer.WriteString(pt.Volume)
	return buffer.String()
}

func GetProductTicker(s string) *ProductTicker {
	resp, err := http.Get(API + "/products/" + s + "/ticker")

	if err != nil {
		log.Println("Error getting product ticker: ", err)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error reading response body: ", err)
		return nil
	}

	productTicker := new(ProductTicker)
	err = json.Unmarshal(body, productTicker)

	if err != nil {
		log.Println("Error decoding body: ", err)
		return nil
	}

	return productTicker
}
