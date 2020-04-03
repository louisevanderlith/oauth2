package core

import "github.com/louisevanderlith/husk"

type Scope struct {
	Name string
	DisplayName string
	Secret string
	Claims []string
}

func (s Scope) Valid() (bool, error){
	return husk.ValidateStruct(&s)
}