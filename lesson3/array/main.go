package main

import "fmt"

// 配列型
func main() {
	// 要素数が3つ
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素数を省略
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	//値の取り出し インデックス番号を指定
	fmt.Println(arr2[0], arr2[1])

	//値の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	// 配列の要素数
	fmt.Println(len(arr1))
}