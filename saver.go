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

func saveToFile(info *RequestInfo, subdir string, data []byte) {
	savePath := path.Join(rootDir, "res/dst/", subdir, info.Path)
	fmt.Println("savePath: ", savePath)
	if pos := strings.LastIndex(savePath, "/"); pos >= 0 {
		fmt.Println("MkdirAll: ", savePath[:pos])
		os.MkdirAll(savePath[:pos], 0644)
	}

	// suffix := contentTypes[info.Type]
	// if !strings.HasSuffix(savePath, suffix) {
	// 	savePath += suffix
	// }
	if err := ioutil.WriteFile(savePath, data, 0644); err != nil {
		fmt.Printf("savePath %s error: %v", savePath, err)
	}
}
