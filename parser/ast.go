package parser

var count int

type Node interface {
	NodeId() int
	AppendChild(Node)
	Front() Node
	Next() Node
}

type BasicNode struct {
	BasicNodeId int
	currChild   int
	children    []Node
}

func (b *BasicNode) AppendChild(n Node) {
	b.children = append(b.children, n)
}

func (b *BasicNode) NodeId() int {
	return b.BasicNodeId
}

func (b *BasicNode) Front() Node {
	b.currChild = 0
	if len(b.children) < 1 {
		return nil
	}
	return b.children[b.currChild]
}

func (b *BasicNode) Next() Node {
	b.currChild = b.currChild + 1
	if b.currChild >= len(b.children) {
		return nil
	}
	return b.children[b.currChild]
}

func CreateBasicNode(id int) BasicNode {
	b := BasicNode{BasicNodeId: id}
	b.children = make([]Node, 0)
	return b
}

type ComponentNode struct {
	BasicNode
	Component  string
	Identifier string
}

type ListNode struct {
	BasicNode
}

type ProgramNode struct {
	BasicNode
}

type AliasNode struct {
	BasicNode
	Alias string
}

func NewComponentNode(l, r Node, comp, identifier string) Node {
	b := CreateBasicNode(count)

	if l != nil {
		b.AppendChild(l)
	}

	if r != nil {
		b.AppendChild(r)
	}

	cn := &ComponentNode{
		BasicNode:  b,
		Component:  comp,
		Identifier: identifier,
	}
	count++
	return cn
}

func NewListNode(l, r Node) Node {
	b := CreateBasicNode(count)

	if l != nil {
		b.AppendChild(l)
	}

	if r != nil {
		b.AppendChild(r)
	}

	ln := &ListNode{
		BasicNode: b,
	}
	count++
	return ln
}

func NewAliasNode(n Node, alias string) Node {
	b := CreateBasicNode(count)
	b.AppendChild(n)

	a := &AliasNode{
		BasicNode: b,
		Alias:     alias,
	}
	count++
	return a
}

func NewProgramNode(n Node) Node {
	b := CreateBasicNode(count)
	b.AppendChild(n)

	e := &ProgramNode{
		BasicNode: b,
	}
	count++
	return e
}
