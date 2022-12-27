package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0]) // 1.1
	// 1.2
	for i, v := range os.Args[1:] {
		fmt.Printf("value = %s, index = %d\n", v, i)
	}
}
