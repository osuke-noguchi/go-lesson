package main

import "fmt"

// interface & nil
func main() {
	// あらゆる型と互換性がある 初期値はnil 演算には使用しない
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 1.1
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

}