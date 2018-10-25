package main

import (
	"Basic/tree"
	"fmt"
)

func main() {
	n := createTree()
	traversingTree(n)

}

func createTree() *tree.Node{
	var root  = &tree.Node{}
	root.Init()
	var t = &tree.Node{
		Name:"牛皮拉拉",
	}
	root.Left = t
	root.Right = &tree.Node{
		Name:"牛皮拉拉右",
	}
	root.Left.Left = &tree.Node{
		Name:"飞上了天",
	}
	root.Right.Left = &tree.Node{
		Name:"地表无敌",
	}
	root.Right.Left.Right = &tree.Node{
		Name:"悟空，师傅被妖怪偷走了。",
	}
	//fmt.Println(root)
	return root
}

func traversingTree(n *tree.Node){
	//中序遍历
	mediumOrder(n)
	fmt.Println()
	//前序遍历
	prologueOrder(n)
}

func mediumOrder(n *tree.Node){
	if n.Left !=nil{
		mediumOrder(n.Left)
	}
	fmt.Printf("%s  ",n.Name)
	if n.Right !=nil{
		mediumOrder(n.Right)
	}

}

func prologueOrder(n *tree.Node){
	fmt.Printf("%s  ",n.Name)
	if n.Left !=nil{
		mediumOrder(n.Left)
	}

	if n.Right !=nil{
		mediumOrder(n.Right)
	}

}

func traverseWithChannel()  {
	
}




