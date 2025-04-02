package main

type Base interface {
	Execute()
}

type Component struct {
	leaf []Leaf
}

func (c *Component) AddLeaf(l Leaf) {
	c.leaf = append(c.leaf, l)
}

func (c *Component) Execute() {

}

type Leaf struct {
}

func (l *Leaf) Execute() {

}

func main() {

}
