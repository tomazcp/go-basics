package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello, Playground") // this is what println does

	io.WriteString(os.Stdout, "Hello, Tomás")
	fmt.Println("")
}
