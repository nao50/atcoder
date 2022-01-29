package main

import "fmt"

func solve(h1, h2 int) int {
	return h1 - h2
}

func main() {
	var h1, h2 int
	fmt.Scan(&h1, &h2)

	result := solve(h1, h2)
	fmt.Println(result)
}
