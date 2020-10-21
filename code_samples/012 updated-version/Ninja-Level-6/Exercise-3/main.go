package main

import "fmt"

func main() {
	defer bar()
	foo()

}

func bar() {
	fmt.Println("pepe")
}

func foo() {
	fmt.Println(":pepelaugh")
}
