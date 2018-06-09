package main

import (
	// "fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	rootDir, _ = os.Getwd()
)

func saveToFile(info *RequestInfo, subdir string, data []byte) error {
	savePath := path.Join(rootDir, "res/dst/", subdir, info.Path)
	//fmt.Println("savePath: ", savePath)
	if pos := strings.LastIndex(savePath, "/"); pos >= 0 {
		os.MkdirAll(savePath[:pos], 0644)
		// err := os.MkdirAll(savePath[:pos], 0644)
		// fmt.Println("MkdirAll: ", savePath[:pos], err)
	}

	// suffix := contentTypes[info.Type]
	// if !strings.HasSuffix(savePath, suffix) {
	// 	savePath += suffix
	// }
	err := ioutil.WriteFile(savePath, data, 0644)
	return err
}
