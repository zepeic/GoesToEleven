package main

import "fmt"

func main() {
	v := foo()
	fmt.Println(v())

}

func foo() func() int {
	return func() int {
		return 42
	}
}
