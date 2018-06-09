package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	err := filepath.Walk("./res/src", func(path string, f os.FileInfo, err error) error {
		fmt.Println("path:", path)
		if f == nil {
			fmt.Printf("filepath.Walk() error %v\n", err)
			return err
		}
		if f.IsDir() {
			fmt.Printf("filepath.Walk() error %v\n", err)
			return nil
		}
		newTask(path, wg)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() error %v\n", err)
	}

	//newTask("./res/src/法老娱乐1.txt", wg)

	wg.Wait()
}
