package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello Playground")
	io.WriteString(os.Stdout, "Hello Playground")
}
