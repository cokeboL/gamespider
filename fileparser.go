package main

import (
	"encoding/json"
	"io/ioutil"
)

func parseJsonFile(filename string) ([]RequestInfo, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ret := []RequestInfo{}
	err = json.Unmarshal(data, &ret)

	return ret, err
}
