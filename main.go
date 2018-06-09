package main

import (
	"fmt"
	"sync"
)

func main() {
	//types := map[string]bool{}
	resources, err := parseJsonFile("./res/src/法老娱乐1.txt")
	if err != nil {
		fmt.Println("parseJsonFile error:", err.Error())
		return
	}

	totalNum := 0
	successNum := 0
	mtx := sync.Mutex{}
	wg := &sync.WaitGroup{}

	fmt.Println("--------------------------------------------------------------------------------------------------")
	fmt.Println("start get")
	fmt.Println("--------------------------------------------------------------------------------------------------")

	downLoadInfos := []*RequestInfo{}
	for _, info := range resources {
		if !needDownLoad(&info) {

		} else {
			downLoadInfos = append(downLoadInfos, &info)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------------------")
	for _, info := range downLoadInfos {
		wg.Add(1)
		getResource(info, func(err interface{}) {
			mtx.Lock()
			defer mtx.Unlock()
			defer wg.Done()
			if err == nil {
				successNum++
			}
		})
	}
	fmt.Println("--------------------------------------------------------------------------------------------------")

	wg.Wait()

	fmt.Printf("%d / %d / %d success\n", successNum, len(downLoadInfos), len(resources))
	// fmt.Println("------------------------")
	// for t, _ := range types {
	// 	fmt.Println(t)
	// }
	// fmt.Println("------------------------")
}
