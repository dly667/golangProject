package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	N int
}

func main(){
	n := MyStruct{ 1 }
	// get
	immutable := reflect.ValueOf(n)
	val := immutable.FieldByName("N").Int()
	fmt.Printf("N=%d\n", val) // prints 1

	// set
	mutable := reflect.ValueOf(&n).Elem()
	mutable.FieldByName("N").SetInt(7)
	fmt.Printf("N=%d\n", n.N) // prints 7
}

func GetAllFieldName(stru interface{}) []interface{}{
	if reflect.TypeOf(stru).Kind() == reflect.Ptr{
		// 获取User的所有属性
		var outFieldName []interface{}
		s :=  reflect.TypeOf(stru)
		ss := s.Elem()
		for i := 0; i <ss.NumField(); i++ {
			outFieldName = append(outFieldName, ss.Field(i).Name)
		}
		return outFieldName
	}else {
		return nil
	}

}

