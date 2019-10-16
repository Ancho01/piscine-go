package main

import "github.com/01-edu/z01"

func main() {

	charX := '0'
	for i := 0; i <= 9; i++ {
		z01.PrintRune(charX)
		charX++
	}
	z01.PrintRune(10)
}
