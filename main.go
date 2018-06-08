package main

import (
	"fmt"
)

func main() {
	resources, err := parseJsonFile("./res/src/法老娱乐1.txt")
	if err != nil {
		fmt.Println("parseJsonFile error:", err.Error())
		return
	}
	for i, v := range resources {
		fmt.Printf("resources %04d: %+v", i, v)
	}
}
