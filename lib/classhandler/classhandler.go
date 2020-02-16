package classhandler

import (
	"../../model/classdata"
	"io/ioutil"
	"net/http"
)

func GetClasses(API_URL string) classdata.WVSClass {
	if result, err := http.Get(API_URL + "class"); err != nil {
		panic(err)
	} else {
		response, failure := ioutil.ReadAll(result.Body)
		if failure != nil {
			panic(failure)
		}


		serialized, ser_err := classdata.UnmarshalWVSClass(response)
		if ser_err != nil {
			panic(ser_err)
		}

		return serialized
	}
}

func GetClassesSearch(API_URL string, searchValue string) classdata.WVSClass {
	if result, err := http.Get(API_URL + "class_search?query=" + searchValue); err != nil {
		panic(err)
	} else {
		response, failure := ioutil.ReadAll(result.Body)
		if failure != nil {
			panic(failure)
		}


		serialized, ser_err := classdata.UnmarshalWVSClass(response)
		if ser_err != nil {
			panic(ser_err)
		}

		return serialized
	}
}
