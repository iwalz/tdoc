package table

import "github.com/iwalz/tdoc/elements"

type cell struct {
	component *elements.Component
	border    int
	width     int
	height    int
}
