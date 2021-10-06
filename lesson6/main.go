package main

import "fmt"

// 関数

func Plus(x, y int) int {
	return x + y
}

func Div(x,y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値の変数を宣言
func Double(price int) (result int) {
	result = price * 2
	return
}

//返り値なし
func NoReturn() {
	fmt.Println("No return")
	// return
}

//関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	// 返り値の破棄
	i4, _ := Div(4, 2)
	fmt.Println(i4)

	i5 := Double(100)
	fmt.Println(i5)

	NoReturn()

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}

	i7 := f(1, 3)
	fmt.Println(i7)

	i8 := func(x, y int) int {
		return x + y
	}(3, 4)

	fmt.Println(i8)

	f2 := ReturnFunc()
	f2()
}