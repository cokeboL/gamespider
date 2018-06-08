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

	successNum := 0
	mtx := sync.Mutex{}
	wg := &sync.WaitGroup{}
	for _, v := range resources {
		// if v.Type != "" {
		// 	types[v.Type] = true
		// }
		wg.Add(1)
		getResource(&v, func(err interface{}) {
			//defer fmt.Println("getResource over callback")
			mtx.Lock()
			defer mtx.Unlock()
			defer wg.Done()
			if err == nil {
				successNum++
			}
		})
	}

	wg.Wait()

	fmt.Printf("%d / %d success\n", successNum, len(resources))
	// fmt.Println("------------------------")
	// for t, _ := range types {
	// 	fmt.Println(t)
	// }
	// fmt.Println("------------------------")
}
