package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() []byte {

	url := "http://wx.qlogo.cn/mmhead/Q3auHgzwzM4QbsClOMQYCebTC18YLSFyMygia7ysLTkOatSQGm7Cgow/132"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "761ed2da-f8fc-41f5-ad3c-786ac35d0370")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
