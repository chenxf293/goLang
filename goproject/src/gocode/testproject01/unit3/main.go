package main

import "fmt"

func main() {
	//+号
	//1.正数 2.相加操作 3.字符串拼接
	var n1 int = +10
	fmt.Println(n1)
	var n2 int = 4 + 7
	fmt.Println(n2)
	var sl string = "abc" + "def"
	fmt.Println(sl)

	// / 除号
	fmt.Println(10 / 3)
	fmt.Println(10.0 / 3)

	// % 取模 等价公式： a%b=a-a/b*b
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)

	//++自增操作
	var a int = 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	//++ 自增 加1操作， --自减，减1操作
	//go语言里，++，--操作非常简单，只能单独使用，不能参与到运算中去
	//go语言里，++，--只能在变量的后面，不能写在变量的前面， --a ++a 错误写法

	//赋值运算符
	var n3 int = 10
	fmt.Println(n3)
	var n4 int = (10+20)%3 + 3 - 7 //=右侧的值运算清楚后，再赋值给=的左侧
	fmt.Println(n4)

	var n5 int = 10
	n5 += 20
	fmt.Println(n5)

	//关系运算符
	fmt.Println(5 == 9) //判断左右两侧的值是否相等，相等返回true，不相等返回的是false， ==不是=
	fmt.Println(5 != 9) //判断不等于
	fmt.Println(5 > 9)
	fmt.Println(5 < 9)
	fmt.Println(5 >= 9)
	fmt.Println(5 <= 9)

	//逻辑运算符
	//与逻辑：&&：两个数值/表达式只要有一侧是false，结果一定为false
	//也叫短路与，只要第一个数值/表达式的结果是false，那么后面的表达式等就不用运算了，直接结果就是false
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	//或逻辑：||：两个数值/表达式只要有一侧是true，结果一定为true
	//也叫短路或，只要第一个数值/表达式的结果是true，那么后面的表达式等就不用运算了，直接结果就是true
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	//非逻辑：取相反的结果
	fmt.Println(!true)
	fmt.Println(!false)

	//其他运算符：&，*
	var age int = 18
	fmt.Println("age对应的存储空间的地址为：", &age) //age对应的存储空间的地址为：0xc000016220

	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr这个指针指向的具体数值为：", *ptr)
}
