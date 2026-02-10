package main

import (
	"fmt"
)

type btree struct {
	value int
	lelf  *btree
	right *btree
}

func main() {

	var m *btree
	fmt.Println(m.value)
}
