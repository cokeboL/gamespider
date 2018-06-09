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
	for _, info := range resources {
		// if info.Type != "" {
		// 	types[info.Type] = true
		// }
		tmp := info
		if needDownLoad(&tmp) {
			totalNum++
			wg.Add(1)
			getResource(&tmp, func(err interface{}) {
				//defer fmt.Println("getResource over callback")
				mtx.Lock()
				defer mtx.Unlock()
				defer wg.Done()
				if err == nil {
					successNum++
				}
			})
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------------------")

	wg.Wait()

	fmt.Printf("%d / %d / %d success\n", successNum, totalNum, len(resources))
	// fmt.Println("------------------------")
	// for t, _ := range types {
	// 	fmt.Println(t)
	// }
	// fmt.Println("------------------------")
}
