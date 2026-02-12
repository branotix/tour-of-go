// package main

// import (
// 	"fmt"
// 	"time"
// )

// func callme1() {
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("calling 1 ", i)
// 	}
// }
// func callme2() {
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("calling 2", i)
// 	}
// }
// func main() {
// 	go callme1()
// 	callme2()
// }

// package main

// import "fmt"

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // send sum to c
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c // receive from c

// 	fmt.Println(x, y, x+y)
// }

package main

import "fmt"

func main() {
	nameList := []string{
		"Pulock", "Ashikur", "Ashif", "Monirul", "Pushpo", "Priya", "Sunny",
	}

	for index, value := range nameList {
		fmt.Println("index ", index, " value ", value)
	}

	var x []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(x)
}
