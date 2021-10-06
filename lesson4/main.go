package main

import "fmt"

// 定数
// 頭文字を大文字にすると他のパッケージから呼び出しができるようになる

const Pi = 3.14

const (
	URL = "https://xxx.xp.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

func main() {
	fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)
}