package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Distance(p1 *Point) int {
	return ((p.X-p1.X)*(p.X-p1.X) + (p.Y-p1.Y)*(p.Y-p1.Y))
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)

	p1 := Point{X: x1, Y: y1}
	p2 := Point{X: x2, Y: y2}

	fmt.Println(p1.Distance(&p2))
}
