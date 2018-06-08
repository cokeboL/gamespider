package main

import (
	"encoding/json"
	"io/ioutil"
)

func parseJsonFile(filename string) ([]map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ret := []map[string]interface{}{}
	err = json.Unmarshal(data, &ret)

	return ret, err
}
