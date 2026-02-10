package main

import "fmt"

func main() {
	c1 := 12 + 1i
	c2 := complex(12, 2)
	fmt.Println("complex c1 ", c1, " complex c2", c2)
	var c3 complex64 = complex64(c1 + c2)
	fmt.Println(c3)
}
