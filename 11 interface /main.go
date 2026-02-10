package main

import "fmt"

type Animal interface {
	Speak() string // যে-ই এনিমেল হতে চায়, তাকে অবশ্যই Speak() করতে হবে
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof! Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func MakeItSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	d := Dog{}
	c := Cat{}

	MakeItSpeak(d) // আউটপুট: Woof!
	MakeItSpeak(c) // আউটপুট: Meow!
}
