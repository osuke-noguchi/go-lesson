package main

import "fmt"

// byte(unit8)型
func main() {

	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// HI
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
}