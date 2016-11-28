package component

import "testing"

func TestStacking(t *testing.T) {
	component1 := &Component{}
	component2 := &Component{}
	component1.Add(component2)
}
