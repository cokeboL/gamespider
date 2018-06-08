package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	contentTypes = map[string]string{
		"image/jpeg": ".jpeg",
	}
)

type Req struct {
	Method         string                 `json:"method"`
	HttpVersion    string                 `json:"httpVersion"`
	Ip             string                 `json:"ip"`
	RawHeaderNames map[string]string      `json:"rawHeaderNames"`
	Headers        map[string]interface{} `json:"headers"`
}

type Res struct {
	Ip             string                 `json:"ip"`
	StatusCode     int                    `json:"statusCode"`
	StatusMessage  string                 `json:"statusMessage"`
	RawHeaderNames map[string]string      `json:"rawHeaderNames"`
	Headers        map[string]interface{} `json:"headers"`
}

type RequestInfo struct {
	Method   string `json:"method"`
	Req      `json:"req"`
	Res      `json:"res"`
	HostIp   string `json:"hostIp"`
	Result   int    `json:"result"`
	Type     string `json:"type"`
	Protocol string `json:"protocol"`
	Hostname string `json:"hostname"`
	Path     string `json:"path"`
	Size     int    `json:"size"`
}

func get(url string) []byte {
	//url := "http://wx.qlogo.cn/mmhead/Q3auHgzwzM4QbsClOMQYCebTC18YLSFyMygia7ysLTkOatSQGm7Cgow/132"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "761ed2da-f8fc-41f5-ad3c-786ac35d0370")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return body
}
