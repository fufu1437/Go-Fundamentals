package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n%3 == 0 && n%5 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}

}
