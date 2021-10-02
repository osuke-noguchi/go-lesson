package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())
}

/*
[go run]
Goプラグラムの手軽な実行方法
ファイルパスを渡して実行
$ go run main.go

[go build]
go buildはOPで与えたGoファイルを実行ファイル形式にコンパイルする
$ go build -o main main.go
-o OPを使用して出力する実行ファイルのファイル名を指定できる

ビルドに成功すると、同じディレクトリに実行ファイルが生成される。
*/