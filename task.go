package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
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

func newTask(src string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	resources, err := parseJsonFile(src)
	if err != nil {
		fmt.Println("parseJsonFile error:", err.Error())
		return
	}

	taskWg := &sync.WaitGroup{}

	successNum := 0
	mtx := sync.Mutex{}

	downLoadInfos := []*RequestInfo{}
	ignoreInfos := []*RequestInfo{}
	for _, info := range resources {
		tmp := info
		if needDownLoad(&tmp) {
			downLoadInfos = append(downLoadInfos, &tmp)
			taskWg.Add(1)
		} else {
			ignoreInfos = append(ignoreInfos, &tmp)
		}
	}

	go func() {
		subdir := src
		if pos := strings.LastIndex(src, "."); pos >= 0 {
			subdir = src[:pos]
		}
		if pos := strings.LastIndex(src, "/"); pos >= 0 {
			subdir = src[pos:]
		} else if pos := strings.LastIndex(src, "\\"); pos >= 0 {
			subdir = src[pos:]
		}
		for _, info := range downLoadInfos {
			newSpider(info, subdir, func(err interface{}) {
				mtx.Lock()
				defer mtx.Unlock()
				defer taskWg.Done()
				if err == nil {
					successNum++
				}
			})
		}

	}()
	taskWg.Wait()

	fmt.Printf("%s finish, success: %d, need download: %d, total: %d success\n", src, successNum, len(downLoadInfos), len(resources))
}
