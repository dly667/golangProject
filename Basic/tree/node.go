package tree

type Node struct{
	Name string
	Left *Node
	Right *Node
	Father *Node
}


func (n *Node)Init() {
	*n = Node{
		Name:"root",
		Father:nil,
		Left:nil,
		Right:nil,
	}

}
