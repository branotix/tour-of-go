package main

import (
	"fmt"
)

var m int = 20000

func main() {
	var j int = m
	fmt.Println("Go has strict rules for curly braces!", m)
	j = 233333333333
	fmt.Println("mmmmmmmmmm", j, m)
}
