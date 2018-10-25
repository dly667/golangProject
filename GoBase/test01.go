package GoBase

import (
	"errors"
	"fmt"
)

func Test01(){
	//定义一个名称为people，类型为type的变量
	var people string
	fmt.Println(people)
	//定义多个变量
	var vn1,vn2,vn3 int
	fmt.Println(vn1,vn2,vn3)
	//定义变量并初始化值
	var v0 int = 1
	fmt.Println(v0)
	//同时初始化多个变量
	var vm1,vm2,vm3 = 1,2,3
	fmt.Println(vm1,vm2,vm3)
	//简单定义多个变量
	vj1,vj2,vj3 := 11,22,33
	fmt.Println(vj1,vj2,vj3)
	//:=这个符号直接取代了var和type,但是只能用在函数内部。
	//_会被丢弃
	_,x := 34,35
	fmt.Println(x)
}

func Test02(){
	//常量定义
	const PI float32= 3.141592611
	const PI1 = 3.141592611

	fmt.Println(PI,PI1)


}

var isActive bool //声明全局变量
var enabled, disabled = true, false //忽略类型的声明
func Test04(){
	//布尔值
	var available bool //一般声明
	valid := false
	//available = true
	fmt.Println(available,valid)

}
func Test05(){
	var c complex64 = 5+5i
	//output:(5+5i)
	fmt.Printf("Value is :%v",c)
}

var frenchHello string = ""
var emptyString string = ""
func Test06(){
	no, yes, maybe := "no","yes","maybe"//简短声明，同时声明多个变量
	japaneseHello := "Konichiwa"//同上
	frenchHello = "Bonjour"//常规赋值
	fmt.Println(no,yes,maybe,japaneseHello,frenchHello)
}

func Test07(){
	//``可声明多行字符串
	m := `hello
	world                       
`
	fmt.Print(m)
}

func Test08(){
	//Go 内置一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包
	//errors来处理错误
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}

//分组声明
func Test09(){
	const(
		i = 100
		pi = 3.1415
		prefix = "Go_"
	)


	var(
		i1 = 1
		pi1 float32
		prefix1 string
	)
	var m = 20000
	fmt.Println(i,pi,prefix,i1,pi1,prefix1,m)
}

//iota枚举
func Test10(){
	const (
		x = iota //x==0
		y = iota //y==1
		z = iota
		w  //常量声明省略值时，默认和之前一个值得字面相同。这里隐式地说w=iota，因此w==3。
		//其实上面y和z可同样不用"=iota"
	)
}


//array、slice、map
//array
func Test2_1(){
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n",arr[0])

	arr1 := [...]int{0,1,2,3,6}
	fmt.Println(arr1)

	doubleArray := [2][4]int{[4]int{1,2,3,4,},[4]int{8,8,8,8}}
	fmt.Println(doubleArray)

	doubleArray1 := [2][4]int{{1,2,3,4},{4,4,4,4}}
	fmt.Println(doubleArray1)
}
//slice
func Test2_2(){
	var s  = []int{1,2,3,4,5}
	fmt.Println(s,len(s),cap(s))
	var s1 = s
	fmt.Println(s1,len(s1),cap(s1))
	s[0] = 0
	fmt.Println(s,len(s),cap(s))
	fmt.Println(s1,len(s1),cap(s1))

	var ss = append(s, 66)
	fmt.Println(ss,len(ss),cap(ss))
	fmt.Println(s,len(s),cap(s))

	var arr = [...]int{33,44,55}
	var sl = arr[0:1]
	arr[1] = 66
	fmt.Println(sl)
	fmt.Println(arr)
	var s2 = append(sl,99)
	fmt.Println(s2)

	fmt.Println(arr)
}

func Test2_3()  {
	//声明式1
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1
	fmt.Println(numbers)

	//声明式2
	//1
	var numbers1 = make(map[string]int)
	//2
	numbers2 := make(map[string]int)
	numbers1["two"] = 2
	fmt.Println(numbers1,numbers2)


	//直接初始化1
	rating := map[string]float32{"C":5}
	//直接初始化2
	var name  =  map[string]string{"A":"liyong"}
	fmt.Println(rating)
	fmt.Println(name)
}