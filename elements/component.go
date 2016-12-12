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
	SetX(int)
	SetY(int)
	X() int
	Y() int
	Added(bool)
	IsAdded() bool
}

type DefaultElement struct {
	index int
	stack []Element
	root  Element
	x     int
	y     int
	added bool
}

func (d *DefaultElement) SetX(x int) {
	d.x = x
}

func (d *DefaultElement) Added(a bool) {
	d.added = a
}

func (d *DefaultElement) IsAdded() bool {
	return d.added
}

func (d *DefaultElement) SetY(y int) {
	d.y = y
}

func (d *DefaultElement) X() int {
	return d.x
}

func (d *DefaultElement) Y() int {
	return d.y
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
