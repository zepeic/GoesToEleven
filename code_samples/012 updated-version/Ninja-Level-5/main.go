package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
	}{
		first: "pepe",
		last:  "hands",
	}
	fmt.Println(p1)
}
