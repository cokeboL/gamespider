package main

import (
	"fmt"
)

func main() {
	types := map[string]bool{}
	resources, err := parseJsonFile("./res/src/法老娱乐1.txt")
	if err != nil {
		fmt.Println("parseJsonFile error:", err.Error())
		return
	}
	for _, v := range resources {
		if v.Type != "" {
			types[v.Type] = true
		}
	}
	fmt.Println("------------------------")
	for t, _ := range types {
		fmt.Println(t)
	}
	fmt.Println("------------------------")
}
