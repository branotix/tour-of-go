package main

import "fmt"

func main() {
	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go mysum(myslice[:5], c)
	go mysum(myslice[:10], c)

	datapaici, datapaini := <-c, <-c
	fmt.Println("data paici ", datapaici, "data pai nai ", datapaini)
}

func mysum(s []int, c chan int) {
	sum := 0
	for _, val := range s {
		sum += val
	}
	c <- sum
}
