package core

import (
	"github.com/louisevanderlith/husk"
	"strings"
)

type profileFilter func(obj Profile) bool

func (f profileFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Profile))
}

//byID filter will filter by client_id
func byTitleID(id string) profileFilter {
	return func(obj Profile) bool {
		return strings.ToLower(obj.Title) == id
	}
}
