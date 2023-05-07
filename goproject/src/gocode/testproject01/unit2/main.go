package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

// 全局变量：定义在函数外的变量
var n7 = 100
var n8 = 9.9

// 设计者认为上面的全局变量的写法太麻烦了，可以一次声明：
var (
	n9  = 500
	n10 = "GO"
)

func main() {
	//定义在{}中的变量叫 局部变量
	//1.变量声明
	var age int
	//2.变量的赋值
	age = 18
	//3.变量的使用
	fmt.Println("age = ", age)

	//变量的四种使用方式
	//第一种：变量的使用方式：指定变量的类型，并且赋值
	var num int = 18
	fmt.Println(num)
	//第二种：指定变量的类型，但是不赋值，使用默认值
	var num2 int
	fmt.Println(num2)
	//第三种： 如果没有写变量的类型，那么根据=后面的值进行判定变量的类型（自动类型推断）
	var num3 = "Tom"
	fmt.Println(num3)
	//第四种：省略var，注意 := 不能写为 =
	num4 := "男"
	fmt.Println(num4)

	//声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4, name, n5 = 10, "Jack", 7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	n6, height := 6.9, 100.6
	fmt.Println(n6)
	fmt.Println(height)

	//输出全局变量
	fmt.Println(n7)
	fmt.Println(n8)
	fmt.Println(n9)
	fmt.Println(n10)
	//Printf函数的作用就是：格式化的，把n10的类型填充到%T的位置上
	fmt.Printf("n10的类型是：%T\n", n10)
	fmt.Println(unsafe.Sizeof(n10))

	//定义浮点类型的数据
	var fnum float32 = 3.14
	fmt.Println(fnum)
	//可以表示正浮点数，也可以表示负的浮点数
	var fnum2 float32 = -3.14
	fmt.Println(fnum2)
	//浮点数可以用于十进制表示形式，也可以用科学计数法表示形式 E 大小写都可以
	var fnum3 float32 = 314e-2
	fmt.Println(fnum3)
	var fnum4 float32 = 314e+2
	fmt.Println(fnum4)
	var fnum5 float32 = 314e-2
	fmt.Println(fnum5)
	var fnum6 float64 = 314e+2
	fmt.Println(fnum6)

	//浮点数可能会有精度的损失，所以通常情况下，建议使用：float64
	var fnum7 float32 = 256.000000916
	fmt.Println(fnum7)
	var fnum8 float64 = 256.000000916
	fmt.Println(fnum8)

	//golang中默认的浮点类型为：float64
	var fnum9 = 3.17
	fmt.Printf("fnum9对应的默认的类型为：%T", fnum9)
}
