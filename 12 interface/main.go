// package main

// import (
// 	"fmt"
// 	"math"
// )

// type vertex struct {
// 	X, Y float64
// }

// func abs(a vertex) float64 {
// 	return math.Sqrt(a.X*a.X + a.Y*a.Y)
// }
// func (v vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
// func main() {
// 	myvertext := vertex{4.0, 5.0}

// 	fmt.Println(abs(myvertext))
// 	fmt.Println(myvertext.Abs())
// }

package main

import "fmt"

func main() {
	x := 10
	y := &x
	fmt.Println(*y)
}
