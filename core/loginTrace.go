package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/secure/core/tracetype"
)

type LoginTrace struct {
	Location  string `hsk:"null;size(128)"`
	IP        string `hsk:"null;size(50)"`
	Allowed   bool   `hsk:"default(true)"`
	ClientID  string `hsk:"size(20)"`
	TraceType tracetype.Enum
}

func (o LoginTrace) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func getRegistrationTrace(r Registration) LoginTrace {
	return LoginTrace{
		Allowed:   true,
		ClientID:  r.ClientID,
		TraceType: tracetype.Register,
	}
}

func getLoginTrace(passed bool) LoginTrace {
	trace := tracetype.Login

	if !passed {
		trace = tracetype.Fail
	}

	return LoginTrace{
		Allowed:   passed,
		//ClientID:  r.ClientID,
		TraceType: trace,
	}
}
