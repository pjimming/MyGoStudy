package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func ffor() {
	mp := make(map[int]float32)
	mp[1] = 1.0
	mp[2] = 2.0
	mp[3] = 3.0

	for k, v := range mp {
		fmt.Printf("k = %d, v = %f\n", k, v)
	}
}

func main() {
	fmt.Println("Hello World!")
	// ffor()
	a, b := swap("aaa", "bbb")
	fmt.Printf("a = %s, b = %s\n", a, b)
}
