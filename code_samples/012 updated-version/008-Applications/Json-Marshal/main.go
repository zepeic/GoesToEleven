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

func main() {
	b1 := bot{
		Shopify: false,
		Fs:      true,
		Ys:      true,
	}
	b2 := bot{
		Shopify: true,
		Fs:      true,
		Ys:      false,
	}
	bots := []bot{
		b1,
		b2,
	}
	bs, err := json.Marshal(bots)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
