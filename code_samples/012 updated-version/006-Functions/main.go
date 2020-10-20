package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	f := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	s1 := foo()
	fmt.Println(s1)
	f(1984)
	s2 := bar()
	fmt.Printf("%T\n", s2)
	i := s2()
	fmt.Println(i)
}

func foo() string {
	s := "hello world"
	return s

}

func bar() func() int {
	return func() int {
		return 451
	}
}
