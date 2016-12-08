package renderer

import (
	"testing"

	"github.com/iwalz/tdoc/elements"
	"github.com/stretchr/testify/assert"
)

func TestAddSingleComponent(t *testing.T) {
	b := &BaseMatrix{rows: 1, columns: 1}
	m := elements.NewMatrix(nil)

	m.Add(elements.NewComponent(nil, nil, "node", "foo", "bar"))

	b.scan(m)
	assert.Equal(t, 1, b.rows)
	assert.Equal(t, 1, b.columns)

	components := b.components
	assert.Equal(t, 1, components[0].X())
	assert.Equal(t, 1, components[0].Y())
}

func TestThreeComponents(t *testing.T) {
	b := &BaseMatrix{rows: 1, columns: 1}
	m := elements.NewMatrix(nil)

	m.Add(elements.NewComponent(nil, nil, "node", "foo", "bar"))
	m.Add(elements.NewComponent(nil, nil, "node", "foo1", "bar1"))
	m.Add(elements.NewComponent(nil, nil, "node", "foo2", "bar2"))

	assert.Equal(t, true, m.HasChilds())

	b.scan(m)
	assert.Equal(t, 2, b.rows)
	assert.Equal(t, 2, b.columns)

	components := b.components
	assert.Equal(t, 1, components[0].X())
	assert.Equal(t, 1, components[0].Y())

	/*assert.Equal(t, 100, components[1].x)
	assert.Equal(t, 0, components[1].y)

	assert.Equal(t, 100, components[2].x)
	assert.Equal(t, 100, components[2].y)*/
}

func TestNestedComponents(t *testing.T) {
	b := &BaseMatrix{rows: 1, columns: 1}
	m := elements.NewMatrix(nil)

	e1 := elements.NewComponent(nil, nil, "node", "foo1", "bar1")
	e2 := elements.NewComponent(nil, nil, "node", "foo2", "bar2")

	e1.Add(elements.NewComponent(nil, nil, "cloud", "foo1", "bar1"))
	e1.Add(elements.NewComponent(nil, nil, "cloud", "foo2", "bar2"))

	e2.Add(elements.NewComponent(nil, nil, "actor", "foo1", "bar1"))
	e2.Add(elements.NewComponent(nil, nil, "actor", "foo2", "bar2"))

	m.Add(e1)
	m.Add(e2)
	assert.Equal(t, true, m.HasChilds())
	assert.Equal(t, true, e1.HasChilds())
	assert.Equal(t, true, e2.HasChilds())

	b.scan(m)
	assert.Equal(t, 2, b.rows)
	assert.Equal(t, 3, b.columns)

	assert.Equal(t, 50, b.heightoffset)
	assert.Equal(t, 50, b.widthoffset)

	components := b.components
	// First component is nested, so we have a border
	assert.Equal(t, 1, components[0].X())
	assert.Equal(t, 1, components[0].Y())

	/*assert.Equal(t, 125, components[1].x)
	assert.Equal(t, 25, components[1].y)

		assert.Equal(t, 150, components[2].x)
		assert.Equal(t, 150, components[2].y)

		assert.Equal(t, 150, components[3].x)
		assert.Equal(t, 150, components[3].y)*/
}

func TestMultiNestedComponents(t *testing.T) {
	b := &BaseMatrix{rows: 1, columns: 1}
	m := elements.NewMatrix(nil)

	e1 := elements.NewComponent(nil, nil, "node", "foo1", "bar1")
	e2 := elements.NewComponent(nil, nil, "node", "foo2", "bar2")
	c1 := elements.NewComponent(nil, nil, "cloud", "foo1", "bar1")
	c1.Add(elements.NewComponent(nil, nil, "cloud", "foo2", "bar2"))
	c1.Add(elements.NewComponent(nil, nil, "cloud", "foo3", "bar3"))
	c1.Add(elements.NewComponent(nil, nil, "cloud", "foo4", "bar4"))
	e1.Add(c1)

	e2.Add(elements.NewComponent(nil, nil, "actor", "foo1", "bar1"))
	e2.Add(elements.NewComponent(nil, nil, "actor", "foo2", "bar2"))

	m.Add(e1)
	m.Add(e2)
	assert.Equal(t, true, m.HasChilds())
	assert.Equal(t, true, e1.HasChilds())
	assert.Equal(t, true, e2.HasChilds())

	b.scan(m)
	assert.Equal(t, 3, b.rows)
	assert.Equal(t, 3, b.columns)
	assert.Equal(t, 75, b.heightoffset)
	assert.Equal(t, 75, b.widthoffset)
}
