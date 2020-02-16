package main

import (
	"./lib/classhandler"
	"./lib/dsbhandler"
	"./lib/feedhandler"
	"fmt"
	"github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
	"net/http"
	"os"
	"strings"
	"time"
)

var API_FEED_URL = os.Getenv("API_FEED_URL")
var API_DSB_URL = os.Getenv("API_DSB_URL")
var API_CLASS_URL = os.Getenv("API_CLASS_URL")
var PORT = os.Getenv("PORT")

func feed(w http.ResponseWriter, r *http.Request) {
	var searchValue = strings.TrimPrefix(r.URL.Path, "/wvs_proxy/feed/")

	feed := feedhandler.GetFeedResults(API_FEED_URL, searchValue)

	result, err := feed.Marshal()
	if err != nil {
		panic(err)
	}

	w.Write(result)
}

func dsb(w http.ResponseWriter, r *http.Request) {
	dsb := dsbhandler.GetDSBResults(API_DSB_URL)

	class := r.URL.Query().Get("class")
	if class != "" {
		modified := dsbhandler.SearchResult(dsb, class)
		result, err := modified.Marshal()
		if err != nil {
			panic(err)
		}

		w.Write(result)
	} else {
		result, err := dsb.Marshal()
		if err != nil {
			panic(err)
		}

		w.Write(result)
	}
}

func dsbSoon(w http.ResponseWriter, r *http.Request) {
	dsb := dsbhandler.GetDSBNextResults(API_DSB_URL)

	result, err := dsb.Marshal()
	if err != nil {
		panic(err)
	}

	w.Write(result)
}

func class(w http.ResponseWriter, r *http.Request) {
	class := classhandler.GetClasses(API_CLASS_URL)

	result, err := class.Marshal()
	if err != nil {
		panic(err)
	}

	w.Write(result)
}

func classSearch(w http.ResponseWriter, r *http.Request) {
	class := classhandler.GetClassesSearch(API_CLASS_URL, r.URL.Query().Get("query"))

	result, err := class.Marshal()
	if err != nil {
		panic(err)
	}

	w.Write(result)
}

func checkEnv() {
	if API_FEED_URL == "" {
		API_FEED_URL = "https://api.thepublictransport.de/wvs_rss/"
	}

	if API_CLASS_URL == "" {
		API_CLASS_URL = "https://api.thepublictransport.de/wvs_school_class/"
	}

	if API_DSB_URL == "" {
		API_DSB_URL = "https://api.thepublictransport.de/dsb/"
	}

	if PORT == "" {
		PORT = "5024"
	}
}

func main() {
	// Checking env
	checkEnv()

	memcached, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LFU),
		memory.AdapterWithCapacity(10000000),
	)
	if err != nil {
		fmt.Println(err)
	}

	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(memcached),
		cache.ClientWithTTL(10 * time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		fmt.Println(err)
	}

	feedHandler := http.HandlerFunc(feed)
	dsbHandler := http.HandlerFunc(dsb)
	dsbNextHandler := http.HandlerFunc(dsbSoon)
	classHandler := http.HandlerFunc(class)
	classSearchHandler := http.HandlerFunc(classSearch)

	http.Handle("/wvs_proxy/feed/", cacheClient.Middleware(feedHandler))
	http.Handle("/wvs_proxy/info/now", cacheClient.Middleware(dsbHandler))
	http.Handle("/wvs_proxy/info/soon", cacheClient.Middleware(dsbNextHandler))
	http.Handle("/wvs_proxy/class", cacheClient.Middleware(classHandler))
	http.Handle("/wvs_proxy/class/search", cacheClient.Middleware(classSearchHandler))

	if err := http.ListenAndServe(":" + PORT, nil); err != nil {
		panic(err)
	}
}
