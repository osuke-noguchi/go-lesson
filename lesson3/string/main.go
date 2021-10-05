package main

import "fmt"

// 文字列型
func main() {

	var s string = "Hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Printf("%T\n", si)

	fmt.Printf(`test
	test
	test`)

	fmt.Printf("\"")
	fmt.Printf(`"`)
		// バイト型
	fmt.Println(s[0])

	fmt.Println(string(s[0]))
}