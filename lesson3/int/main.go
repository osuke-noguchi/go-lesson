package main

import "fmt"

func main() {
	// 数値型
	var i int = 100

	fmt.Println(i + 50)

	// intの型が違う場合は計算できない

	var i2 int64 = 200
	// 型を表示
	fmt.Printf("%T\n", i2)

	//型を変換する
	fmt.Printf("%T\n", int32(i2))
}
