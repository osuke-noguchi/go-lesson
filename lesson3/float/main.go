package main

import "fmt"

// 浮動小数点型
func main() {

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 暗黙的な定義の場合はfloat64になる
	fl := 3.2
	fmt.Println(fl64 + fl)


	// 基本的にfloat32は使わない
	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	//型変換
	fmt.Printf("%T\n", float64(fl32))
}