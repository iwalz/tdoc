package elements

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Stackable interface {
	Add(*Component)
}

type Element interface {
	Add(Element)
	AddRelation(*Relation)
	Relations() []*Relation
	Next() Element
	HasChilds() bool
	Reset()
	Root(Element)
	Added(bool)
	IsAdded() bool
	Remove()
	Stack() []Element
	SetStack([]Element)
}

type DefaultElement struct {
	index     int
	stack     []Element
	relations []*Relation
	added     bool
	root      Element
}

func (d *DefaultElement) Stack() []Element {
	return d.stack
}

func (d *DefaultElement) SetStack(s []Element) {
	d.stack = s
}

func (d *DefaultElement) AddRelation(r *Relation) {
	d.relations = append(d.relations, r)
}

func (d *DefaultElement) Relations() []*Relation {
	return d.relations
}

func (d *DefaultElement) Added(a bool) {
	d.added = a
}

func (d *DefaultElement) IsAdded() bool {
	return d.added
}

func (d *DefaultElement) Remove() {
	index := 0
	found := false
	stack := d.root.Stack()
	for i, v := range stack {
		if v == d {
			index = i
			found = true
			break
		}
	}
	fmt.Println("Blubb")
	if found {
		fmt.Println("Found element on ", index)
		stack = append(stack[:index], stack[index+1:]...)
		spew.Dump(stack)
		d.root.SetStack(stack)
	}
}

func (d *DefaultElement) Add(e Element) {
	d.stack = append(d.stack, e)
	e.Root(d)
}

func (d *DefaultElement) Root(e Element) {
	d.root = e
}

func (d *DefaultElement) Reset() {
	d.index = 0
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
	d.relations = make([]*Relation, 0)
	return d
}

type Component struct {
	DefaultElement
	Typ        string
	Identifier string
	Alias      string
}

func NewComponent(typ, identifier, alias string) *Component {
	d := NewDefaultElement()
	c := &Component{
		DefaultElement: d,
		Typ:            typ,
		Identifier:     identifier,
		Alias:          alias,
	}

	return c
}
