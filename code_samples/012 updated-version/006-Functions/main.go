package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}
type human interface {
	speak()
}

// func (r reciever) identifier(paremeters) (return(s)) {code....}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)

}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},

		ltk: true,
	}
	p1 := person{
		first: "Dr.",
		last:  "No",
	}
	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(sa2)
	sa2.speak()
	fmt.Println(p1)
}
