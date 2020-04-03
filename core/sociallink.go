package core

import "github.com/louisevanderlith/husk"

type SocialLink struct {
	Icon string `hsk:"size(25)"`
	URL  string `hsk:"size(128)"`
}

func (o SocialLink) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
