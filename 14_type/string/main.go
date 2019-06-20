package main

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	// [72 101 108 108 111 44 32 112 108 97 121 103 114 111 117 110 100]
	// 72 -> H, 101 -> e, etc.
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
