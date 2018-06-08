package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	downMtx     = sync.Mutex{}
	downLoading = map[string]bool{}

	contentTypes = map[string]string{
		"text/html":                ".html",
		"text/css":                 ".css",
		"application/javascript":   ".js",
		"image/png":                ".png",
		"application/json":         ".js",
		"application/x-javascript": ".js",
		"image/jpeg":               ".jpeg",
		// "application/x-font-woff":  ".fnt",
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
	Url      string `json:"url"`
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

func checkAndSetDownLoading(url string) bool {
	downMtx.Lock()
	defer downMtx.Unlock()
	_, ok := downLoading[url]
	downLoading[url] = true
	return ok
}

func getResource(info *RequestInfo, cb func(err interface{})) {
	if _, ok := contentTypes[info.Type]; ok && !checkAndSetDownLoading(info.Url) {
		//url := "http://wx.qlogo.cn/mmhead/Q3auHgzwzM4QbsClOMQYCebTC18YLSFyMygia7ysLTkOatSQGm7Cgow/132"
		//fmt.Println("start getResource: ", info.Url)
		go func() {
			defer func() {
				err := recover()
				cb(err)
				//fmt.Println("getResource error: ", err)
			}()

			for i := 0; i < 5; i++ {
				req, err := http.NewRequest("GET", info.Url, nil)
				if err != nil {
					fmt.Println("start getResource error 111: ", err)
					continue
				}
				req.Header.Add("Cache-Control", "no-cache")
				req.Header.Add("Postman-Token", "761ed2da-f8fc-41f5-ad3c-786ac35d0370")

				client := http.Client{
					Timeout: 10 * time.Second,
				}
				res, err := client.Do(req)
				if err != nil {
					fmt.Println("start getResource error 222: ", err)
					continue
				}

				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println("start getResource error 333: ", err)
					continue
				}

				info.Path = req.URL.Path
				//fmt.Println("start saveToFile:", req.URL.Path)
				saveToFile(info, body)
				break
			}
		}()
	} else {
		defer cb("duplication")
	}
}
