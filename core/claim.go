package core

type Claims []Claim

type Claim struct {
	Name  string
	Value interface{}
}