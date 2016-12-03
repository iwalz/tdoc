package elements

type Stackable interface {
	Add(*Component)
}

type Element interface {
	Add(Element)
	Next() Element
}

type DefaultElement struct {
	index int
	stack []Element
}

func (d *DefaultElement) Add(e Element) {
	d.stack = append(d.stack, e)
}

func (d *DefaultElement) Next() Element {
	index := d.index

	if len(d.stack) >= d.index {
		d.index++
		return d.stack[index]
	}

	return nil
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

func NewComponent(l, r Element, typ, identifier string) Element {
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

type List struct {
	DefaultElement
}

func NewList(e1, e2 Element) Element {
	d := NewDefaultElement()
	d.Add(e1)
	d.Add(e2)

	l := &List{
		DefaultElement: d,
	}

	return l
}
