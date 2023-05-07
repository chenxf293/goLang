package main
import "fmt"
func main() {
	//方式1
	//定义map变量
	var a map[int]string
	//只声明map内存是没有分配空间
	//必须通过make函数进行初始化，才会分配空间
	a = make(map[int]string,10)//map可以存放10个键值对
	//将键值对存入map中
	a[1] = "张三"
	a[2] = "里斯"
	a[3] = "王五"
	//输出集合
	fmt.Println(a)

	//方式2
	b := make(map[int]string)
	b[1] = "张三"
	b[2] = "里斯"
	fmt.Println(b)

	//方式3
	c := map[int]string{
		1 : "张三",
		2 : "里斯",
	}
	//增加
	c[3] = "王五"
	fmt.Println(c)

	//删除
	delete(c,3)
	fmt.Println(c)

	//查找：
	value,flag := c[3]
	fmt.Println(value)
	fmt.Println(flag)

	//获取长度
	fmt.Println(len(a))

	//遍历：
	for k,v := range b {
		fmt.Printf("key为：%v value为 %v]\t",k,v)
	}
	fmt.Println("--------------------------------")

	//加深难度
	d := make(map[string]map[int]string)
	//赋值
	d["class1"] = make(map[int]string,3)
	d["class1"][1] = "Jhon"
	d["class1"][2] = "TOm" 
	d["class1"][3] = "Tim"

	d["class2"] = make(map[int]string,3)
	d["class2"][1] = "Jhon2"
	d["class2"][2] = "TOm2" 
	d["class2"][3] = "Tim2"

	for k1,v1 := range d {
		fmt.Println(k1)
		for k2,v2 := range v1 {
			fmt.Printf("学生学号为：%v 学生姓名为 %v \t",k2,v2)
		}
		fmt.Println()
	}
}