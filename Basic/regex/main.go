package main

import (
	"fmt"
	"regexp"
)

func main() {
	const s = `hahah sld  dly667@163.com.cn! sdf hahah sld  bb@162.com@sdfsdf! sdf
hahah sld  sdf@11.com@sdfsdf! sdf
hahah sld  z@163.com@sdfsdf! sdf
hahah sld  aa@163.com@sdfsdf! sdf
`
	re:= regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9.]+[a-zA-Z0-9]+")
	match := re.FindAll([]byte(s),-1)
	fmt.Printf("%s",match)

}
