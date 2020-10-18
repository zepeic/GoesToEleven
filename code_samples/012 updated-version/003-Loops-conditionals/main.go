package main

import "fmt"

func main() {
	switch "Bond" {
	case "Moneypenny":
		fmt.Println("This should not print")
	case "Bond":
		fmt.Println("This should  print2")
	case "Q":
		fmt.Println("prints")
		fallthrough

	default:
		fmt.Println("this is defaults")
	}
}
