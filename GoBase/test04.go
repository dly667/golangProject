package GoBase

import (
	"fmt"
	"reflect"
)

//反射
func ReflectTest()  {
	var x float32= 3.5
	v := reflect.ValueOf(x)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	fmt.Println(v.Float())
}

//channel通信
func sum(a []int, c chan int){
	total := 0
	for _,v:= range a{
		total += v
	}
	c <- total //send total to c
}

func Test04Main()  {
	a := []int{7,2,8,-9,4,0}

	c := make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	x,y := <-c,<-c
	fmt.Println(x,y,x+y)
}
