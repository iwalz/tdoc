package elements

import (
	"errors"
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
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

func IsRelation(r string) (*Relation, bool) {
	firstChar := r[0]
	relation := &Relation{}
	relationString := ""
	var relationChar byte
	var nextChar byte
	index := 0
	fmt.Println("Index first:", index)
	fmt.Println("r:", r)
	fmt.Println("First char:", string(firstChar))
	fmt.Println("Last char:", string(r[len(r)-1]))
	// Both relations set
	if firstChar == '<' && r[len(r)-1] == '>' {

		fmt.Println("Both matched")
		relation.arrowLocation = Both

		if strings.HasPrefix(r, "<|") {
			fmt.Println("Fillede arrpw")
			relation.arrowTypeLeft = Filled
			index = index + 2
			nextChar = r[index]
		} else {
			relation.arrowTypeLeft = Regular
			index = index + 1
			nextChar = r[index]
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
			nextChar = r[index]
			relation.arrowTypeLeft = Filled
		} else {
			index = index + 1
			nextChar = r[index]
			relation.arrowTypeLeft = Regular
		}
		relation.direction = Left
	}
	// Right relation set
	if firstChar != '<' && r[len(r)-1] == '>' {
		relation.arrowLocation = Right
		fmt.Println("Right matched")
		if strings.HasSuffix(r, "|>") {
			relation.arrowTypeRight = Filled
		} else {
			relation.arrowTypeRight = Regular
		}
		relation.direction = Right
	}
	nextChar = r[index]
	fmt.Println("Next char", string(nextChar))
	fmt.Println(index)
	// Set direction
	if nextChar == 'u' {
		relation.direction = Up
		index = index + 1
	}
	if nextChar == 'd' {
		relation.direction = Down
		index = index + 1
	}
	if nextChar == 'l' {
		relation.direction = Left
		index = index + 1
	}
	if nextChar == 'r' {
		relation.direction = Right
		index = index + 1
	}

	if relation.direction == 0 {
		relation.direction = Right
	}
	nextChar = r[index]

	// Default to right, but don't use next character
	/*if relation.direction == 0 {
		relation.direction = Right
	}*/
	fmt.Println("Next char", string(nextChar))
	// Set line type
	if nextChar == ' ' {
		relation.kind = Empty
		relationChar = ' '
	}
	if nextChar == '.' {
		relation.kind = Dotted
		relationChar = '.'
	}
	if nextChar == '=' {
		relation.kind = Thick
		relationChar = '='
	}
	if nextChar == '-' {
		relation.kind = Dashed
		relationChar = '-'
	}

	openIndex := strings.Index(r, "[")
	if openIndex != -1 {
		closeIndex := strings.LastIndex(r, "]")
		if closeIndex == -1 {
			fmt.Println("False in close index")
			return nil, false
		}
		relationString = r[0:openIndex]
		relationString = relationString + r[closeIndex+1:len(r)]
		relation.text = r[openIndex+1 : closeIndex]
	}
	fmt.Println("Index", index)
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
	fmt.Println("Foo", string(nextChar))
	if nextChar == '[' {
		relationChar = relationString[index]
	} else {
		nextChar = r[index]
	}

	if (nextChar == '|' && checkValidSuffix(r)) || relationString == "|>" {
		relation.kind = Dashed
		relation.size = 1
		relation.arrowLocation = Right
		relation.arrowTypeRight = Filled

		return relation, true
	}

	if relation.kind == 0 && !checkValidSuffix(r) {
		fmt.Println("False in kind and prefix")
		return nil, false
	}

	// Allowed characters at this point are:
	// The kind character & >, <, |
	for i := 0; i < len(relationString); i++ {
		if relationString[i] != relationChar && relationString[i] != '>' && relationString[i] != '<' && relationString[i] != '|' {
			fmt.Println(string(relationChar))
			spew.Dump(relationString)
			fmt.Println("False in character check")
			return nil, false
		}
	}
	fmt.Println("Count", relationString)
	count := strings.Count(relationString, string(relationChar))
	if count <= 0 {
		relation.size = 1
	} else {
		relation.size = count
	}

	return relation, true
}

func checkValidSuffix(s string) bool {
	if strings.HasSuffix(s, ">") {
		return true
	}

	if strings.HasSuffix(s, "-") {
		return true
	}

	if strings.HasSuffix(s, " ") {
		return true
	}

	if strings.HasSuffix(s, ".") {
		return true
	}

	if strings.HasSuffix(s, "|>") {
		return true
	}

	return false
}

func (b *Relation) To(e Element) {
	b.element = e
}
