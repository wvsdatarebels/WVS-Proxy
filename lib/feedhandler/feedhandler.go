package feedhandler

import (
	"io/ioutil"
	"../../model/feeddata"
	"net/http"
)

func GetFeedResults(API_URL string, searchValue string) feeddata.FeedData {
	if result, err := http.Get(API_URL + searchValue); err != nil {
		panic(err)
	} else {
		response, failure := ioutil.ReadAll(result.Body)
		if failure != nil {
			panic(failure)
		}


		serialized, ser_err := feeddata.UnmarshalFeedData(response)
		if ser_err != nil {
			panic(ser_err)
		}

		return serialized
	}
}
