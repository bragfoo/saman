package util

import (
	"net/http"
	"io/ioutil"
)

const idsURL = "http://utils.instreet.cc:1273/id"

func GetObjectId() (string) {
	r, _ := http.Get(idsURL)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	return string(body)
}
