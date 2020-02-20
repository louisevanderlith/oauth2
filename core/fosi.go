package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/ory/fosite"
)

type Fosi struct {
	ID string
	fosite.Requester
}

func (o Fosi) Valid() (bool, error){
	return husk.ValidateStruct(&o)
}