// ম্যপের যত সব পাঁচ লাগানো বিষয় রে
//agei bole dibo kon doroner map hobe

package main

import "fmt"

type Studnet map[string]int

func main() {
	var StudnetsList Studnet = Studnet{
		"Pulock Kumar":   2304060,
		"Ashikur Rahman": 2304059,
		"Pushpo Apa":     2304061,
	}
	fmt.Println(StudnetsList)
}

//ভাই ভাই নিচে আরেকটা type এর ম্যাজিক দেখানো হবে তুমি প্রস্তুত তো ম্যজিক দেখার জন্য
