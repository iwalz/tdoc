package elements

import "errors"

type Registry struct {
	identifier map[string]Element
}

func NewRegistry() *Registry {
	i := make(map[string]Element)
	registry := &Registry{
		i,
	}

	return registry
}

func (r *Registry) Add(e Element) error {
	if e == nil {
		return errors.New("Nil given")
	}

	c := e.(*Component)
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

func Get(r *Registry, i string) Element {
	return r.identifier[i]
}
