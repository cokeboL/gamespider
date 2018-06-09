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
	fmt.Println("newTask: ", src)
	defer wg.Done()

	resources, err := parseJsonFile(src)
	if err != nil {
		fmt.Println("parseJsonFile %s error: %v", src, err)
		return
	}

	taskWg := &sync.WaitGroup{}

	successNum := 0
	mtx := sync.Mutex{}

	downLoadInfos := []*RequestInfo{}
	ignoreInfos := []*RequestInfo{}
	for _, info := range resources {
		tmp := info
		if needDownLoad(src, &tmp) {
			downLoadInfos = append(downLoadInfos, &tmp)
			taskWg.Add(1)
		} else {
			ignoreInfos = append(ignoreInfos, &tmp)
		}
	}

	success := []*RequestInfo{}
	failed := []*RequestInfo{}
	go func() {
		subdir := src
		// if pos := strings.LastIndex(subdir, "."); pos >= 0 {
		// 	subdir = src[:pos]
		// }
		if pos := strings.LastIndex(subdir, "/"); pos >= 0 {
			subdir = src[pos:]
		} else if pos := strings.LastIndex(subdir, "\\"); pos >= 0 {
			subdir = src[pos:]
		}
		for _, info := range downLoadInfos {
			tmp := info
			newSpider(tmp, subdir, func(err interface{}) {
				mtx.Lock()
				defer mtx.Unlock()
				defer taskWg.Done()
				if err == nil {
					success = append(success, tmp)
					successNum++
				} else {
					failed = append(failed, tmp)
				}
			})
		}
	}()
	taskWg.Wait()

	sep := "-------------------------------------------------------------------------------------------\n"
	str := "\n" + sep + fmt.Sprintf("task \"%s\" finish, success: %d, need download: %d, total: %d \n\nsuccess: %d\n", src, successNum, len(downLoadInfos), len(resources), successNum)
	for _, info := range success {
		str += ("  " + info.Url + "\n")
	}
	str += fmt.Sprintf("\nfailed: %d\n", len(downLoadInfos)-successNum)
	for _, info := range failed {
		str += ("  " + info.Url + "\n")
	}
	str += sep
	fmt.Println(str)
}
