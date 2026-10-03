package main

import "fmt"

func main() {
	var first, second int
	fmt.Scan(&first)
	fmt.Scan(&second)
	// Print one greeting per line, first name first.
	fmt.Println(first + second)
	fmt.Println(first - second)
	fmt.Println(first == second)
}
