package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1,len2 := len(nums1),len(nums2)

	zzz := len1+len2
	//fmt.Println(zzz)

	// var newArr [zzz]int
	newArr := make([]int,zzz)

	counter := 0
	num2_conter_temp := -1
	if len1 !=0 && len2!=0{
		for _,v := range nums1{
			fmt.Println(v)
			for m,n := range nums2{
				if (m<=num2_conter_temp){
					continue
				}
				if v>n{
					fmt.Println(v,n)
					newArr[counter] = n
					num2_conter_temp = m
					counter++

				}else{
					newArr[counter] = v
					counter++
					//num2_conter_temp = m -1
					break
				}
			}
		}
		fmt.Println(newArr)
		if num2_conter_temp <len(nums2)-1{
			newArr = append(newArr[:counter],nums2[num2_conter_temp+1:]...)
		}else if num2_conter_temp ==len(nums2)-1 {
			newArr = append(newArr[:counter],nums1[len1-(zzz-counter):]...)
		}
	}else {
		newArr = append(nums1,nums2...)
	}


	centerNum := 0.0
	len3 := len(newArr)
	fmt.Println(newArr)
	if len3%2 == 1{
	    centerNum = float64(newArr[len3/2])
	}else{
		t := newArr[len3/2]+newArr[len3/2-1]
	    centerNum = float64(t)/2
	}
	return float64(centerNum)
}

func main(){
	//m :=[]int{1,3,4,5,6,7,8,32,55,66}
	//n :=[]int{2,17,34,56,88,111}
	//m :=[]int{3,4,5}
	//n :=[]int{1,2}
	//m := []int{1,3}
	//n := []int{2}
	m := []int{1,2}
	n := []int{-1,3}
	findMedianSortedArrays(m,n)
}
