// ম্যপের যত সব পাঁচ লাগানো বিষয় রে
//agei bole dibo kon doroner map hobe

// 01 map type of type
// package main

// import "fmt"

// type Studnet map[string]int

// func main() {
// 	var StudnetsList Studnet = Studnet{
// 		"Pulock Kumar":   2304060,
// 		"Ashikur Rahman": 2304059,
// 		"Pushpo Apa":     2304061,
// 	}
// 	fmt.Println(StudnetsList)
// }

//ভাই ভাই নিচে আরেকটা type এর ম্যাজিক দেখানো হবে তুমি প্রস্তুত তো ম্যজিক দেখার জন্য
//02 data type of type

// package main

// import "fmt"

// func main() {
// 	type myint int64

// 	var number myint = 10234014312431
// 	fmt.Println(number)
// }

//03 functional type

// package main

// import "fmt"

// type functionaltype func(string) int

// func main() {
// 	var myfunc functionaltype = func(s string) int {
// 		fmt.Println("hello i am funtional type ", s)
// 		return 1
// 	}

// 	x := func() int {
// 		return 1
// 	}
// 	fmt.Println("just test my func ", x)
// 	fmt.Println(myfunc("hi am i audiable"))
// }

//04 method of struct

// package main

// import "fmt"

// type car struct {
// 	partname string
// 	model    int
// 	price    float32
// }

// func (c car) getdata(pr string, m int, p float32) {
// 	c.partname = pr
// 	c.model = m
// 	c.price = p
// }
// func (c car) showdata() {
// 	fmt.Println("partname ", c.partname, " model ", c.model, " price ", c.price)
// }
// func main() {
// 	var hunda = car{
// 		partname: "jonto machine",
// 		model:    1080,
// 		price:    120.0,
// 	}
// 	hunda.showdata()
// }

//05 আপনি যদি কোনো ফাংশনের ভেতরে আপনার স্ট্রাক্টের ডাটা আপডেট করতে চান, তবে পয়েন্টার ছাড়া সেটা সম্ভব নয়।

// package main

// import "fmt"

// type Studnet struct {
// 	name string
// 	id   int
// }

// func (s *Studnet) getdata(n string, i int) {
// 	s.name = n
// 	s.id = i
// }

// func main() {
// 	pulock := Studnet{
// 		name: "Pulock Kumar",
// 		id:   2304060,
// 	}

// 	pulock.getdata("Suzon Sarkar", 0)
// 	fmt.Println(pulock.name, " ", pulock.id)
// }

//06 check nill value

// package main

// import "fmt"

// type Studnet struct {
// 	name string
// 	id   int
// }

// func (s *Studnet) getdata(n string, i int) {
// 	s.name = n
// 	s.id = i
// }

// func main() {
// 	var s *Studnet
// 	fmt.Println(s == nil)
// }

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// ইনসার্ট করার ফাংশন
func (n *Node) Insert(value int) {
	if value <= n.Data {
		// ছোট বা সমান হলে বামে যাও
		if n.Left == nil {
			n.Left = &Node{Data: value}
		} else {
			n.Left.Insert(value) // বামের নোডকে বলো আবার চেক করতে
		}
	} else {
		// বড় হলে ডানে যাও
		if n.Right == nil {
			n.Right = &Node{Data: value}
		} else {
			n.Right.Insert(value) // ডানের নোডকে বলো আবার চেক করতে
		}
	}
}

// ট্রি ঠিকমতো সাজানো হয়েছে কি না দেখতে প্রিন্ট ফাংশন (In-order Traversal)
func (n *Node) PrintAll() {
	if n == nil {
		return
	}
	n.Left.PrintAll() // আগে ছোটগুলো
	fmt.Printf("%d ", n.Data)
	n.Right.PrintAll() // তারপর বড়গুলো
}

func main() {
	// রুট নোড শুরু করলাম
	tree := &Node{Data: 50}

	// ইচ্ছেমতো ডাটা পুশ করেন
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	fmt.Print("সাজানো ট্রি: ")
	tree.PrintAll()
	// আউটপুট আসবে: 20 30 40 50 60 70 80 (সব সিরিয়ালে!)
}
