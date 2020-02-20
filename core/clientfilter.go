package core

import (
	"github.com/louisevanderlith/husk"
)

type clientFilter func(obj Client) bool

func (f clientFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Client))
}

//byID filter will filter by client_id
func byClientID(id string) clientFilter {
	return func(obj Client) bool {
		return obj.ID == id
	}
}