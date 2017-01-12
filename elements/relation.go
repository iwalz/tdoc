package elements

const (
	Dashed = iota
	Dotted
	Thick
	Empty
)

const (
	Right = iota
	Left
	Up
	Down
)

const (
	Regular = iota
	Both
)

type Relation struct {
	element   Element
	kind      int
	direction int
	size      int
	arrow     int
	text      string
}

func NewRelation(r string) Relation {
	return IsRelation(r)
}

func IsRelation(r string) Relation {
	bytes := []byte(r)
	relation := Relation{}

	for k, v := range bytes {
		if k == 0 && v == 'u' {
			relation.direction = Up
		}
		if k == 0 && v == 'r' {
			relation.direction = Right
		}
		if k == 0 && v == 'l' {
			relation.direction = Left
		}
	}

	return relation
}

func (b *Relation) To(e Element) {
	b.element = e
}
