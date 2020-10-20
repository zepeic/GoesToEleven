package main

import "fmt"

type mytype int

var y int
var x mytype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

}
