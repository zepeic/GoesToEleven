package main

import (
	"encoding/json"
	"fmt"
)

type bot struct {
	Shopify bool
	Fs      bool
	Ys      bool
}

type bots []struct {
	Shopify bool `json:"Shopify"`
	Fs      bool `json:"Fs"`
	Ys      bool `json:"Ys"`
}

func main() {
	s := `[{"Shopify":false,"Fs":true,"Ys":true},{"Shopify":true,"Fs":true,"Ys":false}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	// bots := []bot{}
	var bots []bot

	err := json.Unmarshal(bs, &bots)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range bots {

		fmt.Println("\nBot Number:", i)
		fmt.Println(v.Shopify, v.Fs, v.Ys)
	}
}
