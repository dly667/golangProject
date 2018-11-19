package GoBase

import "fmt"
//struct类型
func Test003(){
	//struct
	type person struct {
		name string
		age int
	}

	//var P = person{"Tom",33}
	//var P = person{name:"Tom",age:44}
	//var P = new(person)
	var P person
	P.name = "bob"
	//P.age = 25
	fmt.Print(P)

}