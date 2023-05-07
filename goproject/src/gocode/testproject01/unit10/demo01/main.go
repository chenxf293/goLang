package main
import "fmt"
func main(){
	//实现的功能：给出五个学生的成绩，求出成绩的总和，平均数
	//给出五个学生的成绩：----> 数组储存
	//定义一个数组
	var scores [5]int
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21
	//求和
	//定义一个变量专门接收成绩的和：
	sum := 0
	for i := 0;i < len(scores);i++ {
		sum += scores[i]
	}
	//平均数
	avg := sum / len(scores)
	//输出
	fmt.Printf("成绩的总和为：%v，成绩的平均数为：%v\n",sum,avg)
	
	var arr [3]int16
	fmt.Println(len(arr))//获取长度
	fmt.Println(arr)//[0 0 0]
	fmt.Printf("arr的地址为：%p\n",&arr)//证明arr中存储的是地址
	//第一个空间的地址
	fmt.Printf("arr的地址为：%p\n",&arr[0])
	//第二个空间的地址
	fmt.Printf("arr的地址为：%p\n",&arr[1])
	//第三个空间的地址
	fmt.Printf("arr的地址为：%p\n",&arr[2])

	//将成绩存入数组
	var scores2 [5]int
	for i := 0;i < len(scores2);i++ {
		fmt.Printf("请录入第%d学生的成绩",i + 1)
		fmt.Scanln(&scores2[i])
	}
	//展示一下班级的每个学生的成绩：（数组进行遍历）
	for i := 0;i < len(scores2);i++ {
		fmt.Printf("第%d个学生的成绩为：%d\n",i+1,scores2[i])
	}
	fmt.Println("--------------------------------")
	for key,val := range scores2 {
		fmt.Printf("第%d个学生的成绩为：%d\n",key+1,val)
	}
	fmt.Println("--------------------------------")
	for _,val := range scores2 {
		fmt.Printf("学生的成绩为：%d\n",val)
	}

	//数组定义初始化
	//第一种
	var arr1 [3]int = [3]int{3,6,9}
	fmt.Println(arr1)
	//第二种
	var arr2 = [3]int{4,5,6}
	fmt.Println(arr2)
	//第三种,长度不指定
	var arr3 = [...]int{4,5,6,8}
	fmt.Println(arr3)
	//第四种,指定索引赋值 索引：数值
	var arr4 = [...]int{2:66,0:33,1:99,3:88}
	fmt.Println(arr4)
	fmt.Printf("%T",arr4)


	//二维数组
	var arr5 [2][3]int16
	fmt.Println(arr)

	fmt.Printf("arr5的地址是：%p\n",&arr5)
	fmt.Printf("arr5[0]的地址是：%p\n",&arr5[0])
	fmt.Printf("arr5[0][0]的地址是：%p\n",&arr5[0][0])

	fmt.Printf("arr5[1]的地址是：%p\n",&arr5[1])
	fmt.Printf("arr5[1][0]的地址是：%p\n",&arr5[1][0])

	//赋值：
	arr5[0][1] = 47
	arr5[0][0] = 82
	arr5[1][1] = 25
	fmt.Println(arr5)

	//初始化操作：
	var arr6 [2][3]int = [2][3]int{{1,4,7},{2,5,8}}
	fmt.Println(arr6)

	//数组遍历
	var arr7 [3][3]int = [3][3]int{{1,4,7},{2,5,8},{3,6,9}}
	fmt.Println(arr7)
	fmt.Println("-----------------------")
	//方式1：普通for循环：
	for i := 0;i < len(arr7);i++ {
		for j := 0;j < len(arr7[i]);j++ {
			fmt.Print(arr7[i][j],"\t")
		}
		fmt.Println()
	}
	fmt.Println("---------------------------")
	//方式2：for range循环：
	for key,value := range arr7 {
		for k,v := range value {
			fmt.Printf("arr[%v][%v]=%v\t",key,k,v)
		}
		fmt.Println()
	}
}