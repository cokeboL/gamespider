package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	rootDir, _ = os.Getwd()
)

func saveToFile(info *RequestInfo, data []byte) {
	savePath := path.Join(rootDir, "res/dst/", info.Path)
	fmt.Println("savePath: ", savePath)
	if pos := strings.LastIndex(savePath, "/"); pos >= 0 {
		fmt.Println("MkdirAll: ", savePath[:pos])
		os.MkdirAll(savePath[:pos], 0644)
	}
	if err := ioutil.WriteFile(savePath, data, 0644); err != nil {
		fmt.Printf("savePath %s error: %v", savePath, err)
	}

}
