package core

import (
	"github.com/louisevanderlith/husk"
)

type fosiFilter func(obj Fosi) bool

func (f fosiFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Fosi))
}

//id filter
func byID(id string) fosiFilter {
	return func(obj Fosi) bool {
		return obj.ID == id
	}
}
