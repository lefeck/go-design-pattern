package main

import "fmt"

const (
	LeafNode = iota
	CompositeNode
)

/*
组件 （Component） 接口描述了树中简单项目和复杂项目所共有的操作。
叶节点 （Leaf） 是树的基本结构， 它不包含子项目, 一般情况下， 叶节点最终会完成大部分的实际工作， 因为它们无法将工作指派给其他部分。
容器 （Container）——又名 “组合 （Composite）”——是包含叶节点或其他容器等子项目的单位。
容器不知道其子项目所属的具体类， 它只通过通用的组件接口与其子项目交互。
容器接收到请求后会将工作分配给自己的子项目， 处理中间结果， 然后将最终结果返回给客户端。
*/

type Component interface {
	Parent() Component
	SetParent(component Component)
	Name() string
	SetName(name string)
	AddChild(component Component)
	Print(pre string)
}

type component struct {
	name   string
	parent Component
}

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(component Component) {
	c.parent = component
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {}

func (c *component) Print(string) {}

type Leafs struct {
	component
}

func NewLeaf() *Leafs {
	return &Leafs{}
}

func (l *Leafs) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, l.name)
}

type Composites struct {
	component
	childs []Component
}

func NewComposite() *Composites {
	return &Composites{
		childs: make([]Component, 0),
	}
}

func (c *Composites) AddChild(component Component) {
	component.SetParent(c)
	c.childs = append(c.childs, component)
}

func (c *Composites) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}

func main() {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print("")
}
