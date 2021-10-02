package main

import "fmt"

var i5 int = 500

func 	outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string  = "Golang"
	)
	fmt.Println(i2, s2)

	// 変数のみ定義
	// 宣言した変数を使用しない場合はエラーになる
	var i3 int //intの初期値は0
	var s3 string //stringの初期値は''(空文字)
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 値の更新
	i = 150
	fmt.Println(i)

	// 暗黙的な定義　型指定の必要がない,関数内のみで使用可能
	// 基本的には明示的に型指定を行う方が良い
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 = "Hello"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()
}
