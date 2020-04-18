package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(operazioni(2, 5))
}

func operazioni(x, y int) (s, p int) {
	return x + y, x * y
}
