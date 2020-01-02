package core

import (
	"github.com/louisevanderlith/husk"
)

type userFilter func(obj User) bool

func (f userFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(User))
}

//Email filter will filter by email and verification status
func emailFilter(email string) userFilter {
	return func(obj User) bool {
		return obj.Email == email && obj.Verified
	}
}
