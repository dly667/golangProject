package GoBase

import "fmt"

//goto
func Test3_1(){
	i := 0
	Here:
		println(i)
		goto Here
}

//for
func Test3_2(){
	sum := 0;
	for index:=0;index<10;index++{
		sum += index;
	}
	fmt.Println(sum)
}

//whie
func Test3_3(){
	sum := 1
	for sum < 1000{
		sum += sum
	}
	fmt.Println(sum)
}

//for range
func Test3_4(){
	var abc = map[string]string{"A":"C","b":"d"}
	for k := range abc{
		fmt.Println(k)
	}
}

//switch
func Test3_5(){
	i := 3
	switch i {
	case 3:
		fmt.Println(11)
		fallthrough
	case 1, 2:
		fmt.Println(22)

	}
}

//func Golang函数
func Function()  string {
	test := "'test"
	return test
}
//如果声明了返回值变量，则可以直接return
func Function1() (m int,n int){
	//不需要用:=
	m = 111
	n = 222


	return
}
//函数变参 arg 是一个slice
func Function2(arg ...int){
	for k,v := range arg{
		fmt.Println(k,v)
	}
}

//函数 传指针
func add(a *int) int{
	*a = *a+1
	return *a
}

//函数 defer
//如果有很多调用defer，那么defer是采用后进先出模式，所以如下代码会输出4 3 2 1 0
func deferTest(){
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
//函数作为值、类型
type testInt func(int) bool //声明一个函数类型
func isOdd(integer int) bool{
	if integer%2 == 0{
		return false
	}
	return true
}
func filter(f testInt){
	println(f(9))
}
func Test001(){
	//x := 1
	//r := add(&x)
	//fmt.Println(r)
	//fmt.Println(x)

	//deferTest()

	filter(isOdd)

}