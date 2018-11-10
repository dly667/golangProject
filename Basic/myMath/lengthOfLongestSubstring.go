package myMath

import (
	"fmt"
	"unicode/utf8"
)

func lengthOfLongestSubstring(s string) int{
	var start  = 0
	var maxLength  =0
	var lastOccurred  = make(map[byte]int)

	for k,v := range []byte(s){
		if value,ok := lastOccurred[v];ok && value >=start{

			start = value+1

		}
		if k - start +1 >maxLength{
			maxLength = k - start +1
		}
		lastOccurred[v] = k
		//fmt.Println(k,v)
		//fmt.Println(start,maxLength,lastOccurred)
	}
	return maxLength

}

func getChineseString(s string) int{
	temp := []byte(s)
	temp1:= []rune(s)
	fmt.Println(temp)
	fmt.Println(temp1)
	fmt.Println(utf8.RuneCount(temp))
	fmt.Println(s)
	return 0
}

func main()  {

	a := lengthOfLongestSubstring("abacsad")
	fmt.Println(a)

	getChineseString("我是哈哈哈！!！!")
}
