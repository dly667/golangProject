package myMath

import "fmt"

func Sqrt(x float64) float64{
	z := 0.0
	for i := 0;i<1000;i++{
		z -= (z*z - x) / (2*x)
	}
	return z
}

func Test() rune{
	var a rune = 1

	t := false
	if t{
		var arr [2]int
		arr[0] = 1
	}

	return a
}

func Sum() rune{
	sum := 0;
	for index:=0;index<10;index++{
		sum+= index
	}
	fmt.Println("sum is equal to",sum)
	return 0
}

//func RangeMap() rune{
//	for k,v:=range map {
//		fmt.Println("map's key:",k);
//		fmt.Println("map's val:",v);
//	}
//	return 0
//}