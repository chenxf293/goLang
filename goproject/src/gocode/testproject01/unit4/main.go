package main

import "fmt"

func main() {
	//实现功能：如果口罩的库存小于30个，提示：库存不足
	var count int = 100
	//单分支：
	if count < 30 {
		fmt.Println("存量不足")
	}

	//if后面表达式，返回结果一定是true或者false
	//如果返回结果为true的话，那么{}中的代码就会执行
	//如果返回结果为false的话，那么{}中的代码就不会执行
	//if后面一定要有空格，和条件表达式分隔开来
	//{}一定不能省略
	//条件表达式左右的（）是建议省略的
	//在golang里，if后面可以并列的加入变量的定义

	if count2 := 20; count2 < 30 {
		fmt.Println("存量不足")
	}

	//如果口罩的库存小于30个，提示：库存不足，否则：库存充足
	var count2 int = 70
	if count2 < 30 {
		fmt.Println("库存不足")
	} else {
		fmt.Println("库存充足")
	}

	//多分支，如果已经走了一个分支了，那么下面的分支就不会再去判断执行了
	var score int = 81
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	//switch分支
	//switch后面是一个表达式，这个表达式的结果依次跟case进行比较，满足结果的话就执行冒号后面的代码。
	//default是用来“兜底”的一个分支，其他case分支都不走的情况下就会走default分支
	//default分支可以放在任意位置上，不一定非要放在最后
	var score2 int = 81
	switch score2 / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("B")
	case 8:
		fmt.Println("C")
	case 7:
		fmt.Println("D")
	case 6:
		fmt.Println("E")
	case 5:
		fmt.Println("F")
	case 4:
		fmt.Println("G")
	case 3:
		fmt.Println("H")
	case 2:
		fmt.Println("I")
	case 1:
		fmt.Println("J")
	case 0:
		fmt.Println("K")
	default:
		fmt.Println("ERROR!")
	}
}
