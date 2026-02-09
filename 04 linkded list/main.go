package main

import "fmt"

type Student struct {
	name       string
	id         int
	deparmtnet string
	next       *Student
}

func main() {
	var head *Student = nil

	var pulock Student = Student{name: "pulock", id: 23, deparmtnet: "cse"}
	head = &pulock

	var ashikur Student = Student{name: "asik", id: 23, deparmtnet: "cse"}
	pulock.next = &ashikur
	var puspo Student = Student{name: "puspo", id: 23, deparmtnet: "cse"}

	ashikur.next = &puspo

	current := head
	for current != nil {
		fmt.Println(current.name, current.id, current.deparmtnet)
		current = current.next
	}

}
