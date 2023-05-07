package main
import "fmt"

//定义老师结构体，将老师中的各个属性 统一放入结构体中管理：
type Teacher struct{
	//变量名字大写外界可以访问这个属性
	Name string
	Age int
	School string
}
//定义Person结构体
type Person struct{
	Name string
}
//给Person结构体绑定方法test
func (p Person) test(){//参数名字随便起
	fmt.Println(p.Name)
}

func main () {
	//创建老师结构体的实例，对象，变量
	var t1 Teacher //var a int
	fmt.Println(t1)//在未赋值时默认值：{0}
	t1.Name = "xxxx"
	t1.Age = 45
	t1.School = "school"
	fmt.Println(t1)

	//创建结构体对象：
	var p Person
	p.Name = "丽丽"
	p.test()


}