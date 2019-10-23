package main

import "github.com/01-edu/z01"

func PointOne(n *int) {

	*n = 1
}

func main() {
	n := 0
	PointOne(&n)
	z01.Println(n)
}
