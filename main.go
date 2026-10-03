package main

import "fmt"

// square should return n multiplied by itself.
func square(n int) int {
	// TODO: replace the placeholder below with the real result.
	return n * n
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(square(n))
}
