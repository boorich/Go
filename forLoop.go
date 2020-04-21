package main

import (
	"fmt"
)

func main() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 {
			continue
		}
		fmt.Println("... continuing ...")
	}
}
