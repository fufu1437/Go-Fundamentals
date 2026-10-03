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
	name, _ := r.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")
	qtyLine, _ := r.ReadString('\n')
	qty, _ := strconv.Atoi(strings.TrimSpace(qtyLine))
	priceLine, _ := r.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceLine), 64)
	total := float64(qty) * price

	// TODO: replace this Println with one fmt.Printf that lays the
	// four values out in the receipt columns described in the exercise.
	// fmt.Println(name, qty, price, total)
	fmt.Printf("%-12s %3d x %6.2f = %8.2f", name, qty, price, total)
}
