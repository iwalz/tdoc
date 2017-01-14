package elements

import (
	"errors"
	"strings"
)

// kind
const (
	_ = iota // Ignore 0 as it's int default
	Dashed
	Dotted
	Thick
	Empty
)

// direction & arrowLocation (left, right and both)
const (
	Right = iota
	Left
	Up
	Down
	Both
)

// arrowType
const (
	Regular = iota
	Filled
)

type Relation struct {
	element        Element
	kind           int
	direction      int
	size           int
	arrowTypeLeft  int
	arrowTypeRight int
	arrowLocation  int
	text           string
}

func NewRelation(r string) (*Relation, error) {
	relation, ok := IsRelation(r)
	if !ok {
		return nil, errors.New(r + " is not a relation")
	}
	return relation, nil
}

func (relation *Relation) extractText(s string) string {
	if relation.text != "" {
		return s
	}

	var relationString string
	openIndex := strings.Index(s, "[")
	if openIndex != -1 {
		closeIndex := strings.LastIndex(s, "]")
		if closeIndex == -1 {
			return s
		}
		relationString = s[0:openIndex]
		relationString = relationString + s[closeIndex+1:len(s)]
		relation.text = s[openIndex+1 : closeIndex]
	}

	return relationString
}

func (relation *Relation) checkKind(b byte) byte {
	var relationChar byte
	if b == ' ' {
		relation.kind = Empty
		relationChar = ' '
	}
	if b == '.' {
		relation.kind = Dotted
		relationChar = '.'
	}
	if b == '=' {
		relation.kind = Thick
		relationChar = '='
	}
	if b == '-' {
		relation.kind = Dashed
		relationChar = '-'
	}

	return relationChar
}

func (relation *Relation) setDirection(b byte, index int) int {
	if b == 'u' {
		relation.direction = Up
		index = index + 1
	}
	if b == 'd' {
		relation.direction = Down
		index = index + 1
	}
	if b == 'l' {
		relation.direction = Left
		index = index + 1
	}
	if b == 'r' {
		relation.direction = Right
		index = index + 1
	}
	if relation.direction == 0 {
		relation.direction = Right
	}

	return index
}

func (relation *Relation) setArrows(r string, index int) int {
	// Both relations set
	firstChar := r[0]
	if firstChar == '<' && r[len(r)-1] == '>' {
		relation.arrowLocation = Both

		if strings.HasPrefix(r, "<|") {
			relation.arrowTypeLeft = Filled
			index = index + 2
		} else {
			relation.arrowTypeLeft = Regular
			index = index + 1
		}

		if strings.HasSuffix(r, "|>") {
			relation.arrowTypeRight = Filled
		} else {
			relation.arrowTypeRight = Regular
		}
		relation.direction = Right
	}
	// Left relation set
	if firstChar == '<' && r[len(r)-1] != '>' && relation.arrowLocation == 0 {
		relation.arrowLocation = Left
		if strings.HasPrefix(r, "<|") {
			index = index + 2
			relation.arrowTypeLeft = Filled
		} else {
			index = index + 1
			relation.arrowTypeLeft = Regular
		}
		relation.direction = Left
	}
	// Right relation set
	if firstChar != '<' && r[len(r)-1] == '>' {
		relation.arrowLocation = Right
		if strings.HasSuffix(r, "|>") {
			relation.arrowTypeRight = Filled
		} else {
			relation.arrowTypeRight = Regular
		}
		relation.direction = Right
	}

	return index
}

func IsRelation(r string) (*Relation, bool) {
	relation := &Relation{}
	relationString := ""
	var relationChar byte
	var nextChar byte

	index := relation.setArrows(r, 0)
	nextChar = r[index]

	// Set direction
	index = relation.setDirection(nextChar, index)
	nextChar = r[index]

	// Set line type
	if nextChar == '[' {
		r = relation.extractText(r)
		relationString = r
		nextChar = r[index]
		relationChar = nextChar
	}
	relationChar = relation.checkKind(nextChar)

	relationString = relation.extractText(r)
	if relationString == "" {
		relationString = r[index:]
	} else {
		relationString = relationString[index:]
	}

	if nextChar == '>' || relationString == ">" {
		relation.kind = Dashed
		relation.size = 1
		relation.arrowLocation = Right
		relation.arrowTypeRight = Regular

		return relation, true
	}

	if (nextChar == '|' && strings.HasSuffix(r, ">")) || relationString == "|>" {
		relation.kind = Dashed
		relation.size = 1
		relation.arrowLocation = Right
		relation.arrowTypeRight = Filled

		return relation, true
	}

	// Allowed characters at this point are:
	// The kind character & >, <, |
	for i := 0; i < len(relationString); i++ {
		if relationString[i] != relationChar && relationString[i] != '>' && relationString[i] != '<' && relationString[i] != '|' {
			return nil, false
		}
	}
	count := strings.Count(relationString, string(relationChar))
	if count <= 0 {
		relation.size = 1
	} else {
		relation.size = count
	}

	return relation, true
}

func (b *Relation) To(e Element) {
	b.element = e
}
