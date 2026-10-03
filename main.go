package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	// TODO: parse `line` with strconv.Atoi, which hands back two values:
	// the number it read and an error. Work out which of the two tells you
	// whether the parse actually worked, then print the success line or the
	// failure word the task asks for. Replace the placeholder below.
	v, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("bad")
	} else {
		fmt.Println("ok", v)
	}
}
