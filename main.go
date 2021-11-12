package main

import "fmt"

func main() {
	res := sum(1, 2)
	fmt.Printf("res: %d", res)
}

func sum(a, b int) int {
	return a + b
}
