package core

import "github.com/louisevanderlith/husk"

type Profile struct {
	Title        string   `hsk:"size(128)"`
	Description  string   `hsk:"size(512)" json:",omitempty"`
	ContactEmail string   `hsk:"size(128)" json:",omitempty"`
	ContactPhone string   `hsk:"size(20)" json:",omitempty"`
	ImageKey     husk.Key `hsk:"null"`
	Clients      []Client
	SocialLinks  []SocialLink `json:",omitempty"`
	APIKeys      map[string]string
}

func (p Profile) Valid() (bool, error) {
	return husk.ValidateStruct(&p)
}

func (p Profile) GetClaims(requested ...string) []string {
	return nil
}
