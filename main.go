package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func increment(n *int) {
	// TODO: make the int that n points at one larger
	(*n)++
}

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(line))
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		v, _ := strconv.Atoi(p)
		nums = append(nums, v)
	}
	for i := range nums {
		increment(&nums[i])
	}
	for _, v := range nums {
		fmt.Println(v)
	}
}
