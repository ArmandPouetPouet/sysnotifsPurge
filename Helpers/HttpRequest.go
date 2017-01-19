package Helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//SetResponse function with json output
func SetResponse(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if i != nil {
		if err := json.NewEncoder(w).Encode(i); err != nil {
			panic(err)
		}
	}
}

//Get performs a http GET call and read body as byte[]
func Get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		panic(err)
	}
	return body
}

//Delete performs a http DELETE call
func Delete(url string) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
