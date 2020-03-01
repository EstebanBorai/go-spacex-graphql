package util

import (
	"io/ioutil"
	"log"
	"net/http"
)

// BodyString converts response body bytes into string
func BodyString(resp *http.Response, str chan string) {
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	str <- string(bytes)
}
