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
	//fmt.Println("savePath: ", savePath)
	if pos := strings.LastIndex(savePath, "/"); pos >= 0 {
		fmt.Println("MkdirAll: ", savePath[:pos])
		os.MkdirAll(savePath[:pos], 0644)
	}
	// if encoding, ok := info.Req.Headers["accept-encoding"]; ok {
	// 	fmt.Println("savePath 111: ", encoding)
	// 	if enc, ok := encoding.(string); ok {
	// 		if enc == "gzip" {
	// 			fmt.Println("savePath 222: ", enc)
	// 			if decrypted, err := gzipDecode(data); err == nil {
	// 				if err := ioutil.WriteFile(savePath, decrypted, 0644); err != nil {
	// 					fmt.Printf("savePath %s error: %v", savePath, err)
	// 				}
	// 			} else {
	// 				fmt.Printf("savePath %s error: %v", savePath, err)
	// 			}
	// 		}
	// 	}
	// } else {
	// 	if err := ioutil.WriteFile(savePath, data, 0644); err != nil {
	// 		fmt.Printf("savePath %s error: %v", savePath, err)
	// 	}
	// }

	// suffix := contentTypes[info.Type]
	// if !strings.HasSuffix(savePath, suffix) {
	// 	savePath += suffix
	// }
	if err := ioutil.WriteFile(savePath, data, 0644); err != nil {
		fmt.Printf("savePath %s error: %v", savePath, err)
	}
}
