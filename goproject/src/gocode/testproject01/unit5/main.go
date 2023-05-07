package main

import "fmt"

func main() {
	//实现一个功能：求和：1+2+3+4+5
	var sum int = 0
	for i := 1; i < 6; i++ {
		sum += i
	}
	fmt.Println(sum)
	//注意：for的初始表达式 不能用var定义变量的形式，要用 :=
}
