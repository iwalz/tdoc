package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardRelations(t *testing.T) {
	r1, ok := IsRelation("l-->")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r1)
	assert.Equal(t, Left, r1.direction)
	assert.Equal(t, Regular, r1.arrowTypeRight)
	assert.Equal(t, 0, r1.arrowTypeLeft)

	r2, ok := IsRelation("d|>")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r2)
	assert.Equal(t, Down, r2.direction)
	assert.Equal(t, 1, r2.size)
	assert.Equal(t, Filled, r2.arrowTypeRight)
	assert.Equal(t, 0, r2.arrowTypeLeft)
	assert.Equal(t, Right, r2.arrowLocation)

	r3, ok := IsRelation("foo-->")
	assert.Equal(t, false, ok)
	assert.Nil(t, r3)

	r4, ok := IsRelation("u-->")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r4)
	assert.Equal(t, Up, r4.direction)
	assert.Equal(t, 2, r4.size)
	assert.Equal(t, Regular, r4.arrowTypeRight)
	assert.Equal(t, 0, r4.arrowTypeLeft)

	r5, ok := IsRelation("u-f-->")
	assert.Equal(t, false, ok)
	assert.Nil(t, r5)
}

func TestRelationWithText(t *testing.T) {
	r1, ok := IsRelation("<-[This is a ] test]->")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r1)
	assert.Equal(t, "This is a ] test", r1.text)
	assert.Equal(t, 2, r1.size)
	assert.Equal(t, Right, r1.direction)
	assert.Equal(t, Both, r1.arrowLocation)
	assert.Equal(t, Regular, r1.arrowTypeLeft)
	assert.Equal(t, Regular, r1.arrowTypeRight)

	r2, ok := IsRelation("<l-[This is a ] test]->")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r2)
	assert.Equal(t, "This is a ] test", r2.text)
	assert.Equal(t, 2, r2.size)
	assert.Equal(t, Left, r2.direction)
	assert.Equal(t, Both, r2.arrowLocation)
	assert.Equal(t, Regular, r2.arrowTypeLeft)
	assert.Equal(t, Regular, r2.arrowTypeRight)

	r3, ok := IsRelation("<|l-[This is a ] test]->")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r3)
	assert.Equal(t, "This is a ] test", r3.text)
	assert.Equal(t, 2, r3.size)
	assert.Equal(t, Left, r3.direction)
	assert.Equal(t, Both, r3.arrowLocation)
	assert.Equal(t, Filled, r3.arrowTypeLeft)
	assert.Equal(t, Regular, r3.arrowTypeRight)

	r4, ok := IsRelation("<|l-[This is a ] test]-|>")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r4)
	assert.Equal(t, "This is a ] test", r4.text)
	assert.Equal(t, 2, r4.size)
	assert.Equal(t, Left, r4.direction)
	assert.Equal(t, Both, r4.arrowLocation)
	assert.Equal(t, Filled, r4.arrowTypeLeft)
	assert.Equal(t, Filled, r4.arrowTypeRight)
}

func TestLeftRelation(t *testing.T) {
	r1, ok := IsRelation("<.[Foo].")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r1)
	assert.Equal(t, "Foo", r1.text)
	assert.Equal(t, 2, r1.size)
	assert.Equal(t, Left, r1.direction)
	assert.Equal(t, Left, r1.arrowLocation)
	assert.Equal(t, Regular, r1.arrowTypeLeft)
	assert.Equal(t, Dotted, r1.kind)
	assert.Equal(t, 0, r1.arrowTypeRight)

	r2, ok := IsRelation("<u===[Foo ✓ Bar]=|>")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r2)
	assert.Equal(t, "Foo ✓ Bar", r2.text)
	assert.Equal(t, 4, r2.size)
	assert.Equal(t, Up, r2.direction)
	assert.Equal(t, Both, r2.arrowLocation)
	assert.Equal(t, Regular, r2.arrowTypeLeft)
	assert.Equal(t, Thick, r2.kind)
	assert.Equal(t, Filled, r2.arrowTypeRight)

	r3, ok := IsRelation("<|u==[Foo ✓ Bar]=")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r3)
	assert.Equal(t, "Foo ✓ Bar", r3.text)
	assert.Equal(t, 3, r3.size)
	assert.Equal(t, Up, r3.direction)
	assert.Equal(t, Left, r3.arrowLocation)
	assert.Equal(t, Filled, r3.arrowTypeLeft)
	assert.Equal(t, Thick, r3.kind)
	assert.Equal(t, 0, r3.arrowTypeRight)
}

func TestSingleRelations(t *testing.T) {
	r1, ok := IsRelation("r[Foo ✓ Bar]>")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r1)
	assert.Equal(t, "Foo ✓ Bar", r1.text)
	assert.Equal(t, 1, r1.size)
	assert.Equal(t, Right, r1.direction)
	assert.Equal(t, Right, r1.arrowLocation)
	assert.Equal(t, 0, r1.arrowTypeLeft)
	assert.Equal(t, Dashed, r1.kind)
	assert.Equal(t, Regular, r1.arrowTypeRight)

	r2, ok := IsRelation("l[Foo ✓ Bar]|>")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r2)
	assert.Equal(t, "Foo ✓ Bar", r2.text)
	assert.Equal(t, 1, r2.size)
	assert.Equal(t, Left, r2.direction)
	assert.Equal(t, Right, r2.arrowLocation)
	assert.Equal(t, 0, r1.arrowTypeLeft)
	assert.Equal(t, Dashed, r2.kind)
	assert.Equal(t, Filled, r2.arrowTypeRight)
}

func TestRelationWithoutArrows(t *testing.T) {
	r1, ok := IsRelation("u-[Foo ✓ Bar]---")
	assert.Equal(t, true, ok)
	assert.NotNil(t, r1)
	assert.Equal(t, "Foo ✓ Bar", r1.text)
	assert.Equal(t, 4, r1.size)
	assert.Equal(t, Up, r1.direction)
	assert.Equal(t, 0, r1.arrowLocation)
	assert.Equal(t, 0, r1.arrowTypeLeft)
	assert.Equal(t, Dashed, r1.kind)
	assert.Equal(t, 0, r1.arrowTypeRight)
}