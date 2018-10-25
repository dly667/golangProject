package mymath

import "fmt"

//杨辉三角形练习
func ShowYangHuiTriangle(line int) {
	results := make(map[int][]int)


	for i:=0;i<line;i++ {

		temp_arr:=make([]int, i+1, i+1)

		//输出样式空格
		for l:=0;l<(line-i);l++ {
			fmt.Printf(" ")
		}
		//第一层例外处理
		if i == 0{

			fmt.Println(1)

			temp_arr[i] = 1

		}else {
			for j:=0;j<=i;j++{
				//第一列和最后一列都为1
				if j==0 || j==i{
					fmt.Print(1)
					temp_arr[j] = 1
					fmt.Printf(" ")
				}else {
					//通过计算上层的两个数字得到结果
					temp_arr[j] = results[i-1][j-1] + results[i-1][j]
					fmt.Print(results[i-1][j-1] + results[i-1][j])
					fmt.Printf(" ")
				}


			}
			fmt.Print("\r\n")

		}
		results[i] = temp_arr


	}

}
