package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func speak(p person) {
	fmt.Println("I am a person I can speak")
	fmt.Println(p.first)
	fmt.Println(p.last)
	fmt.Println(p.age)
}

func main() {
	p1 := person{
		first: "John",
		last:  "Wu ",
		age:   19,
	}
	speak(p1)
}
