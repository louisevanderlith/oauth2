package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/ory/fosite"
)

type AuthCode struct {
	Active bool `json"active"`
	Code string
	fosite.Requester
}

func (o AuthCode) Valid() (bool, error){
	return husk.ValidateStruct(&o)
}
