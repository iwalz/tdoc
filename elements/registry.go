package elements

import (
	"errors"
	"fmt"
)

type Registry struct {
	identifier map[string]*Component
}

func NewRegistry() *Registry {
	i := make(map[string]*Component)
	registry := &Registry{
		i,
	}

	return registry
}

func (r *Registry) Add(c *Component) error {
	if c == nil {
		return errors.New("Nil given")
	}

	if c.Alias != "" {
		_, ok := r.identifier[c.Alias]
		if ok {
			return errors.New("Alias " + c.Alias + " already exists")
		}

		r.identifier[c.Alias] = c
		return nil
	}

	if c.Identifier != "" {
		_, ok := r.identifier[c.Identifier]
		if ok {
			return errors.New("Identifier " + c.Identifier + " already exists")
		}

		r.identifier[c.Identifier] = c
		return nil
	}

	return errors.New("Neither alias nor identifier set")
}

func Get(r *Registry, i string) *Component {
	c, ok := r.identifier[i]
	fmt.Println(ok)
	if ok {
		return c
	}

	return nil
}
