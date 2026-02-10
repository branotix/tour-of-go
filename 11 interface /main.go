package main

import "fmt"

type Human interface {
	sex() string
}

type AdultBoy struct {
	age             int
	sextime         int
	sallaryPerMonth float64
}

func (b AdultBoy) sex() string {
	return fmt.Sprintf("I am adult boy, my sex time is %v min", b.sextime)
}

type AdultGirl struct {
	age        int
	orgasmtime int
	boobsSize  float64
}

func (g AdultGirl) sex() string {
	return fmt.Sprintf("I am adult girl, my orgasm time is %v min", g.orgasmtime)
}

// এই ফাংশনটা ইন্টারফেস রিসিভ করে সরাসরি স্ট্রিং রিটার্ন করবে
func common(h Human) string {
	return h.sex()
}

func main() {
	// ভ্যালু দিয়ে অবজেক্ট তৈরি
	pulock := AdultBoy{age: 25, sextime: 15}
	anisha := AdultGirl{age: 22, orgasmtime: 10}

	// এখন প্রিন্ট হবে ম্যজিক!
	fmt.Println(common(pulock))
	fmt.Println(common(anisha))
}
