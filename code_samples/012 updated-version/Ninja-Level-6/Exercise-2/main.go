package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := foo(x...)
	fmt.Println(y)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := bar(a)
	fmt.Println(b)

}
func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
