package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	tryTimes = 3
	timeout = time.Second * 10
	tryDelay = time.Duration(time.Second / 10)
	totalDelay = time.Duration(time.Second / 10)

	err := filepath.Walk("./res/src", func(path string, f os.FileInfo, err error) error {
		if f == nil {
			fmt.Printf("walk src dir error %v\n", err)
			return err
		}
		if f.IsDir() {
			//fmt.Printf("walk src dir error %v\n", err)
			return nil
		}
		tmp := path
		wg.Add(1)
		go newTask(tmp, wg)
		return nil
	})
	if err != nil {
		fmt.Printf("walk src dir error %v\n", err)
	}

	//newTask("./res/src/法老娱乐1.txt", wg)

	wg.Wait()
}
