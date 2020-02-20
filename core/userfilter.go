package core

import (
	"github.com/louisevanderlith/husk"
	"log"
)

type userFilter func(obj User) bool

func (f userFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(User))
}

//Email filter will filter by email and verification status
func byEmail(email string) userFilter {
	log.Printf("Looking for %s\r\n", email)
	return func(obj User) bool {
		log.Printf("against %s\r\n", obj.Email)
		return obj.Email == email && obj.Verified
	}
}
