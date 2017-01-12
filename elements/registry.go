package elements

import "errors"

type Registry struct {
	alias      map[string]Element
	identifier map[string]Element
}

func NewRegistry() *Registry {
	a := make(map[string]Element)
	i := make(map[string]Element)
	registry := &Registry{
		alias:      a,
		identifier: i,
	}

	return registry
}

func (r *Registry) Add(e Element) error {
	if e == nil {
		return errors.New("Nil given")
	}

	c := e.(*Component)
	if c.Alias != "" {
		_, ok := r.alias[c.Alias]
		if ok {
			return errors.New("Alias " + c.Alias + " already exists")
		}

		r.alias[c.Alias] = c
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

func GetByAlias(r *Registry, a string) Element {
	return r.alias[a]
}

func GetByIdentifier(r *Registry, i string) Element {
	return r.identifier[i]
}
