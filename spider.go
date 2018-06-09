package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	delay      = time.Second / 100
	totalDelay = time.Duration(0)
)

type Req struct {
	Method         string                 `json:"method"`
	HttpVersion    string                 `json:"httpVersion"`
	Ip             string                 `json:"ip"`
	RawHeaderNames map[string]string      `json:"rawHeaderNames"`
	Headers        map[string]interface{} `json:"headers"`
	RawHeaders     map[string]interface{} `json:"rawHeaders"`
}

type Res struct {
	Ip             string                 `json:"ip"`
	StatusCode     int                    `json:"statusCode"`
	StatusMessage  string                 `json:"statusMessage"`
	RawHeaderNames map[string]string      `json:"rawHeaderNames"`
	Headers        map[string]interface{} `json:"headers"`
	RawHeaders     map[string]interface{} `json:"rawHeaders"`
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

func needDownLoad(info *RequestInfo) bool {
	if _, ok := contentTypes[info.Type]; ok && !checkAndSetDownLoading(info.Url) && !strings.Contains(info.Url, "?") {
		return true
	}
	return false
}
func getResource(info *RequestInfo, cb func(err interface{})) {
	//url := "http://wx.qlogo.cn/mmhead/Q3auHgzwzM4QbsClOMQYCebTC18YLSFyMygia7ysLTkOatSQGm7Cgow/132"
	// if info.Url == "http://wx.qlogo.cn/mmhead/Q3auHgzwzM4QbsClOMQYCebTC18YLSFyMygia7ysLTkOatSQGm7Cgow/132" {
	// 	fmt.Println("------------------------------------------------------------------------------------------")
	// 	fmt.Println("get: ", info.Url)
	// 	fmt.Println("------------------------------------------------------------------------------------------")
	// }
	fmt.Println(info.Url)
	var outerr interface{} = nil

	go func(delaytime time.Duration) {
		defer func() {
			if err := recover(); err != nil {
				cb(err)
			}
			cb(outerr)
			//fmt.Println("getResource error: ", err)
		}()

		time.Sleep(delaytime)

		for i := 0; i < 5; i++ {
			time.Sleep(delay)
			req, err := http.NewRequest("GET", info.Url, nil)
			if err != nil {
				outerr = err
				fmt.Println("getResource error 111: ", err)
				continue
			}

			// for key, khead := range info.Req.RawHeaderNames {
			// 	if vhead, ok := info.Req.Headers[key]; ok {
			// 		if strVHead, ok := vhead.(string); ok {
			// 			req.Header.Add(khead, strVHead)
			// 		}

			// 	}
			// }

			for k, v := range info.Req.RawHeaders {
				if str, ok := v.(string); ok {
					req.Header.Add(k, str)
					fmt.Println("Add Header: ", info.Url, k, v)
				} else {
					if arr, ok := v.([]string); ok {
						for _, str := range arr {
							req.Header.Add(k, str)
							fmt.Println("Add Header: ", info.Url, k, v)
						}
					}
				}

			}

			client := http.Client{
				Timeout: 60 * time.Second,
			}
			res, err := client.Do(req)
			if err != nil {
				outerr = err
				fmt.Println("getResource error 222: ", err)
				continue
			}

			if res.StatusCode != 200 {
				fmt.Println(res.StatusCode, res.Status)
				outerr = res.Status
				fmt.Println("getResource error 333: ", outerr)
				continue
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				outerr = err
				fmt.Println("getResource error 444: ", err)
				continue
			}

			info.Path = req.URL.Path
			//fmt.Println("start saveToFile:", req.URL.Path)
			saveToFile(info, body)
			return
		}
		fmt.Println("getResource error 555: ", outerr)
	}(totalDelay)
	totalDelay += delay
}
