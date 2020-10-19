package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["John_wu"] = []string{"demon", "god", "himeself"}
	delete(m, "no_dr")
	for i, v := range m {
		fmt.Printf("Record: %v\n", i)
		for j, val := range v {
			fmt.Printf("\tindex :%v\tvalue: %v\n", j, val)
		}
	}

}
