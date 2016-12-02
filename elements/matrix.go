package elements

type Matrix struct {
	stack []*Component
}

func (m *Matrix) Add(c *Component) {
	m.stack = append(m.stack, c)
}

func NewMatrix() *Matrix {
	return &Matrix{}
}
