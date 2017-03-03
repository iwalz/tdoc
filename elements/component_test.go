package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedComponents(t *testing.T) {
	c1 := &Component{}
	c2 := &Component{}
	assert.False(t, c1.HasChilds())
	c1.Add(c2)

	assert.Equal(t, c2, c1.Next())
	assert.True(t, c1.HasChilds())
	assert.Nil(t, c1.Next())
	c1.Reset()
	assert.Equal(t, c2, c1.Next())

	c1.Added(true)
	assert.True(t, c1.IsAdded())
}

func TestRelationsComponent(t *testing.T) {
	c1 := &Component{}
	c2 := &Component{}
	r := &Relation{}
	r.To(c2)
	c1.AddRelation(r)

	assert.Equal(t, c1.Relations()[0], r)
	assert.Equal(t, c1.Relations()[0].Element(), c2)
}

func TestRemove(t *testing.T) {
	c1 := NewComponent("foo", "bar", "blubb")
	c2 := NewComponent("foo1", "bar", "blubb")
	c3 := NewComponent("foo2", "bar", "blubb")

	c1.Add(c2)
	c1.Add(c3)
	assert.Equal(t, c2, c1.Stack()[0])
	assert.Equal(t, c3, c1.Stack()[1])
	assert.Equal(t, 2, len(c1.Stack()))

	c2.Remove()
	assert.Equal(t, 1, len(c1.Stack()))
	assert.Equal(t, c3, c1.Stack()[0])
}
