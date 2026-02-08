package main

import "fmt"

func jogkoro(a, b int) int {
	return a + b
}

func multipleret(a, b int) (int, int) {
	return a + b, a * b
}
func main() {

	var a, b int
	fmt.Print("Enter kor a er man ")
	fmt.Scan(&a, &b)
	result := jogkoro(a, b)
	fmt.Println("jog koreci ", result)

	result1, result2 := multipleret(a, b)

	fmt.Println("jog koreci ", result1, " gun koreci : ", result2)
}
