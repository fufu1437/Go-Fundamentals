package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	// TODO: accumulate the sum of 1..n with a loop, and print that instead of the placeholder.

	v := 0
	for i := 1; i <= n; i++ {
		v += i
	}

	fmt.Println(v)
}
