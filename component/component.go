package component

type Stackable interface {
	Add(*Component)
}

type Component struct {
	Typ        string
	Identifier string
	Alias      string
	Stack      []*Component
}

func (c *Component) Add(comp *Component) {
	c.Stack = append(c.Stack, comp)
}
