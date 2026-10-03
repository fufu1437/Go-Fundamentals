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
	line = strings.TrimRight(line, "\r\n")
	line = strings.TrimSpace(line)
	// TODO: print the text of line, uppercased, with no whitespace around it.
	// Right now it prints the line untouched, which is wrong for every test.
	fmt.Println(strings.ToUpper(line))
}
