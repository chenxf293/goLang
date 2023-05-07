package main
import "fmt"
func main(){
	//定义数组：
	var intarr [6]int = [6]int{3,6,9,1,4,7}
	//切片构建在数组之上
	//定义一个切片名字为slice，[]动态变化的数组长度不写，int类型，intarr是原数组
	//[1:3]切片 -  切出的一段片段 - 索引：从1开始，到3结束（不包含3）- ）[1,3)
	//var slice []int = intarr[1:3]
	slice := intarr[1:3]
	//输出数组：
	fmt.Println("intarr",intarr)
	//输出切片
	fmt.Println("intarr",slice)
	//切片元素个数：
	fmt.Println("slice的元素个数",len(slice))
	//获取切片的容量：容量可以动态变化
	fmt.Println("slice的容量",cap(slice))
	fmt.Printf("数组中下标为1位置的地址：%p",&intarr[1])
	fmt.Printf("切片中下标为0位置的地址：%p",&slice[0])
	slice[1] = 16
	fmt.Println("intarr:",intarr)
	fmt.Println("slice:",slice)

	//定义切片：make函数的三个参数：1.切片类型 2.切片长度 3.切片的容量
	slice2 := make([]int,4,20)
	fmt.Println(slice2)
	fmt.Println("切片的长度:",len(slice2))
	fmt.Println("切片的容量:",cap(slice2))
	slice2[0] = 66
	slice2[1] = 88
	fmt.Println(slice2)

	slice3 := []int{1,4,7}
	fmt.Println(slice3)
	fmt.Println("切片的长度:",len(slice3))
	fmt.Println("切片的容量:",cap(slice3))

	slice2[2] = 99
	slice2[3] = 100
	//方式1：普通for循环
	for i := 0;i < len(slice2);i++ {
		fmt.Printf("slice[%v] = %v \t",i,slice2[i] )
	}
	fmt.Println("\n----------------------------")
	//方式2：for-range循环：
	for i, v := range slice2 {
		fmt.Printf("下标：%v，元素：%v\n",i,v)
	}

	//切片动态增长
	//定义数组：
	var intarr2 [6]int = [6]int{1,4,7,3,6,9}
	//定义切片
	var slice4 []int = intarr2[1:4]
	fmt.Println(len(slice))

	slice5 := append(slice4,88,50)
	fmt.Println(slice5)
	fmt.Println(slice4)
	//底层原理：
	//1.底层追加元素的时候对数组进行扩容，老数组扩容为新数组：
	//2.创建一个新数组，将老数组中的4，7，3复制到新数组中，在新数组中追加88，50
	//3.slice5底层数组的指向， 指向的是新数组
	//4.往往我们在使用追加的时候其实想要做的效果给slice4追加
	slice4 = append(slice4,88,50)
	fmt.Println(slice4)
	//5.底层的新数组 不能直接维护 需要通过切片间接维护操作

	slice6 := append(slice4,slice5...)
	fmt.Println(slice6)

	//拷贝
	//定义切片：
	var a []int = []int{1,4,7,3,6,9}
	var b []int = make([]int,10)
	
	copy(b,a)//将a中对应数组中元素内容复制到b中对应的数组中
	fmt.Println(b)

}