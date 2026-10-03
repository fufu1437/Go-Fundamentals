package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	words := strings.Fields(strings.TrimSpace(line))
	counts := make(map[string]int)

	// TODO: put every word of words into counts, then set these two numbers.
	different := 0
	once := 0
	for _, v := range words {
		counts[v]++
	}
	for _, v := range counts {
		if v == 1 {
			once++
		}
		different++
	}
	_ = counts

	fmt.Println(different)
	fmt.Println(once)
}
