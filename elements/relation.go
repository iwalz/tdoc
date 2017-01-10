package elements

const (
	Dashed = iota
	Dotted
	Empty
)

type Relation interface {
	To(Element)
	Kind(int)
	Size(int)
}

type BaseRelation struct {
	element Element
	kind    int
	size    int
}

// Right relation
type RRelation struct {
	BaseRelation
}

// Left relation
type LRelation struct {
	BaseRelation
}

// Down relation
type DRelation struct {
	BaseRelation
}

func (b *BaseRelation) To(e Element) {
	b.element = e
}
