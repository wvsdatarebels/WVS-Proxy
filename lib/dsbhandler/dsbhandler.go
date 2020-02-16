package dsbhandler

import (
	"bytes"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"../../model/dsbdata"
	"../../model/dsbnextdata"
	"net/http"
	"os"
	"time"
)

var DSB_USER = os.Getenv("DSB_USER")
var DSB_PASSWORD = os.Getenv("DSB_PASSWORD")

func GetDSBResults(API_URL string) dsbdata.DSBData {
	var dsbRequestBody = []byte(`{"username":"` + DSB_USER + `", "password": "` + DSB_PASSWORD + `"}`)

	c := cache.New(5 * time.Minute, 10 * time.Minute)

	data, found := c.Get("dsb_response_key")
	if !found {
		req, err := http.NewRequest("POST", API_URL + "today/get", bytes.NewBuffer(dsbRequestBody))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, readError := ioutil.ReadAll(resp.Body)
		if readError != nil {
			panic(readError)
		}

		fmt.Print(string(body))

		c.Set(string(body), "dsb_response_key", cache.DefaultExpiration)

		responseData, parseError := dsbdata.UnmarshalDSBData(body)
		if parseError != nil {
			panic(parseError)
		}

		return responseData
	} else {
		responseData, parseError := dsbdata.UnmarshalDSBData([]byte(data.(string)))
		if parseError != nil {
			panic(parseError)
		}

		return responseData
	}
}

func GetDSBNextResults(API_URL string) dsbnextdata.DSBNextData {
	var dsbRequestBody = []byte(`{"username":"` + DSB_USER + `", "password": "` + DSB_PASSWORD + `"}`)

	c := cache.New(5 * time.Minute, 10 * time.Minute)

	data, found := c.Get("dsb_next_response_key")
	if !found {
		req, err := http.NewRequest("POST", API_URL + "next/get", bytes.NewBuffer(dsbRequestBody))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, readError := ioutil.ReadAll(resp.Body)
		if readError != nil {
			panic(readError)
		}

		fmt.Print(string(body))

		c.Set(string(body), "dsb_next_response_key", cache.DefaultExpiration)

		responseData, parseError := dsbnextdata.UnmarshalDSBNextData(body)
		if parseError != nil {
			panic(parseError)
		}

		return responseData
	} else {
		responseData, parseError := dsbnextdata.UnmarshalDSBNextData([]byte(data.(string)))
		if parseError != nil {
			panic(parseError)
		}

		return responseData
	}
}

func SearchResult(data dsbdata.DSBData, class string) dsbdata.DSBData {
	sliced := data
	var newData []dsbdata.Result

	for _, element := range sliced.Result {
		if element.SchoolClassBefore == class {
			newData = append(newData, element)
		}
	}

	sliced.Result = newData

	return sliced
}
