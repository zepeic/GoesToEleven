package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	if v, ok := m["Barnabas"]; ok {

		fmt.Println(v)
		fmt.Println(ok)
	}
	for k, v := range m {
		fmt.Println(k, v)

	}
	xi := []int{4, 5, 7, 8, 9, 42}
	for i, v := range xi {
		fmt.Println(i, v)
	}

}
