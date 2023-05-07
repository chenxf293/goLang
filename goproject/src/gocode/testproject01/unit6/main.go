package main

import (
	"fmt"
)

// 自定义函数：功能：两个数相加
func cal(num1 int, num2 int) int {
	var sum int = 0
	sum = num1 + num2
	return sum
}

// 自定义函数：多个返回值
func cal2(num1 int, num2 int) (int, int) {
	var sum int = 0
	sum = num1 + num2

	sub := num1 - num2
	return sum, sub
}

// 升级写法
func cal3(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

// 定义一个函数，函数的参数为：可变参数... 参数的数量可变
// args...int 可以传入任意多个数量的int类型的数据 传入0个，1个，，，n个
func test(args ...int) {
	//函数内部处理可变参数的时候，将可变参数当作切片来处理
	//遍历可变参数
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func test2(num int) {
	num = 10
	fmt.Println("test2---", num)
}

func test3(num *int) {
	*num = 30
}

func test4(num1 int, num2 float32, testFunc func(int)) {
	testFunc(10)
	fmt.Println("---test4", num1, num2)
}

func main() {
	//调用函数
	sum := cal(10, 20)
	fmt.Println(sum)

	//多个返回值
	sum1, sub := cal2(10, 20)
	fmt.Println(sum1, sub)

	sum2, _ := cal2(10, 20)
	fmt.Println(sum2)

	//可变参数
	test()
	fmt.Println("-------------")
	test(2)
	fmt.Println("-------------")
	test(33, 44, 55, 66, 77)

	//基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值。
	var num int = 30
	test2(num)
	fmt.Println("main---", num)

	//以值传递方式的数据类型，如果希望在函数内的变量能修改函数外的变量，可以传入变量的地址&。
	//函数内以指针的方式操作变量。从效果来看类似引用传递。
	var num2 int = 10
	fmt.Println(&num2)
	test3(&num2)
	fmt.Println(num2)

	//函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了
	function := test2
	fmt.Printf("function的类型是：%T、test2函数的类型是：%T\n", function, test2)
	function(10)

	//定义一个函数，把另一个函数作为形参
	test4(10, 9.8, test2)
	test4(11, 10.9, function)

	//自定义数据类型：（相当于起别名）：给int类型起了别名叫myInt类型
	type myInt int

	var num3 myInt = 30
	fmt.Println("num1", num3)

	var num4 int = 30
	num4 = int(num3) //虽然是别名，但是在go中编译识别的时候还是认为myInt和int不是一种数据类型
	fmt.Println("num4", num4)

}
