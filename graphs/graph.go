package graphs

type Node struct {
	Id         int
	Distance   int
	Color      string
	Visited    bool
	Neighbours []*Node
}

func NewNode(id int, neighbours []*Node) *Node {
	return &Node{
		Id:         id,
		Visited:    false,
		Neighbours: neighbours,
	}
}
