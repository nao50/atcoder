package main

import "fmt"

func solve(m int) int {
	switch {
	default:
		return m
	case m <= 99:
		return 0
	case 100.0 <= m && m <= 5000.0:
		return int((float32(m) / 1000) * 10)
	case 6000 <= m && m <= 30000:
		return int((float32(m) / 1000) + 50)
	case 35000 <= m && m <= 70000:
		return int((((float32(m) / 1000) - 30) / 5) + 80)
	case 70000 < m:
		return 89
	}
}

func main() {
	var m int
	fmt.Scan(&m)

	result := solve(m)

	fmt.Printf("%02d\n", result)
}
