package elements

type Stackable interface {
	Add(*Component)
}

type Element interface {
	Add(Element)
	Next() Element
	Parent(Element)
	Root() Element
	HasChilds() bool
	Reset()
}

type DefaultElement struct {
	index int
	stack []Element
	root  Element
}

func (d *DefaultElement) Add(e Element) {
	d.stack = append(d.stack, e)
}

func (d *DefaultElement) Parent(p Element) {
	d.root = p
}

func (d *DefaultElement) Reset() {
	d.index = 0
}

func (d *DefaultElement) Root() Element {
	return d.root
}

func (d *DefaultElement) Next() Element {
	index := d.index

	if len(d.stack) > d.index {
		d.index = d.index + 1
		return d.stack[index]
	}

	return nil
}

func (d *DefaultElement) HasChilds() bool {
	if len(d.stack) > 0 {
		return true
	}

	return false
}

func NewDefaultElement() DefaultElement {
	d := DefaultElement{}
	d.stack = make([]Element, 0)
	return d
}

type Component struct {
	DefaultElement
	Typ        string
	Identifier string
	Alias      string
}

func NewComponent(l, r Element, typ, identifier, alias string) Element {
	d := NewDefaultElement()
	if l != nil {
		d.Add(l)
	}
	if r != nil {
		d.Add(r)
	}

	c := &Component{
		DefaultElement: d,
		Typ:            typ,
		Identifier:     identifier,
		Alias:          alias,
	}

	return c
}

type Matrix struct {
	DefaultElement
}

func NewMatrix(e1 Element) Element {
	d := NewDefaultElement()
	if e1 != nil {
		d.Add(e1)
	}

	m := &Matrix{
		DefaultElement: d,
	}

	return m
}
