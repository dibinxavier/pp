package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var endpoint = "http://ip-api.com/batch";
	ip := []string{"103.198.128.77"}
	ipapiClient := http.Client{};
	data, _ := json.Marshal(ip)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(data))
	resp, err := ipapiClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var country []IpApiResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	json.Unmarshal(body, &country)
	fmt.Print(country[0].CountryCode, "\n")
	// country[0].CountryCode = "IN"
	if country[0].CountryCode != "SG" && country[0].CountryCode != "US" && country[0].CountryCode != "IN" {
		fmt.Print("Invalid country");
	} else {
		fmt.Print("Valid country");
	}
}

type IpApiResponse struct {
	CountryCode string `json:"countryCode"`
}

var endpoint = "http://ip-api.com/batch"
ip := []string{ipAddress}
ipapiClient := http.Client{}
data, _ := json.Marshal(ip)
req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(data))
resp, err := ipapiClient.Do(req)
if err != nil {
	logger.Error.Println(err)
}
var country []IpApiResponse
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	logger.Error.Println(err)
}
json.Unmarshal(body, &country)
