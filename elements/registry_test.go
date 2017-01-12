package elements

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistryAddId(t *testing.T) {
	r := NewRegistry()
	c := NewComponent("cloud", "foo", "")

	r.Add(c)
	c1 := GetByIdentifier(r, "foo")
	assert.Equal(t, c, c1)
}

func TestRegistryAddAlias(t *testing.T) {
	r := NewRegistry()
	c := NewComponent("cloud", "foo", "bar")

	r.Add(c)
	c1 := GetByAlias(r, "bar")
	assert.Equal(t, c, c1)

	c2 := GetByIdentifier(r, "foo")
	assert.Equal(t, nil, c2)
}

func TestErrors(t *testing.T) {
	r := NewRegistry()
	c := NewComponent("cloud", "", "")

	err := r.Add(nil)
	assert.Equal(t, errors.New("Nil given"), err)

	err = r.Add(c)
	assert.Equal(t, errors.New("Neither alias nor identifier set"), err)

	c1 := NewComponent("cloud", "foo", "bar")
	err = r.Add(c1)
	assert.Equal(t, nil, err)

	err = r.Add(c1)
	assert.Equal(t, errors.New("Alias bar already exists"), err)

	c2 := NewComponent("cloud", "foo", "")
	err = r.Add(c2)
	assert.Equal(t, nil, err)

	err = r.Add(c2)
	assert.Equal(t, errors.New("Identifier foo already exists"), err)
}
