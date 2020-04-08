package main

import (
	"fmt" // Un package nella libreria standard di Go.
	// Libreria matematica, con alias locale m
	// Sì, un web server!
)

func main() {
	fmt.Println("hello world")
	fmt.Println(operazioni(2, 5))

}

func operazioni(x, y int) (s, p int) {
	return x + y, x * y
}
