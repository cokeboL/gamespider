package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var (
	rootDir, _ = os.Getwd()
)

func saveToFile(info *RequestInfo, data []byte) {
	savePath := path.Join(rootDir, info.Path)
	err := ioutil.WriteFile(savePath, data, 0644)
	fmt.Println(err)
}
