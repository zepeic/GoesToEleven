package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a, b, c)
}
func foo() int {
	return 42
}

func bar() (int, string) {
	return 1942, "Big Brother is watching you"

}
