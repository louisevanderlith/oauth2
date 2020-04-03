package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Profiles husk.Tabler
	Scopes husk.Tabler
	Users husk.Tabler
	Forgotten husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Profiles: husk.NewTable(Profile{}, serials.GobSerial{}),
		Scopes: husk.NewTable(Scope{}, serials.GobSerial{}),
		Users: husk.NewTable(User{}, serials.GobSerial{}),
		Forgotten: husk.NewTable(Forgot{}, serials.GobSerial{}),
	}

	seed()
}

func Shutdown() {
	ctx.Profiles.Save()
	ctx.Users.Save()
	ctx.Forgotten.Save()
}

func seed() {
	err := ctx.Users.Seed("db/users.seed.json")

	if err != nil {
		panic(err)
	}

	err = ctx.Users.Save()

	if err != nil {
		panic(err)
	}

	err = ctx.Profiles.Seed("db/profiles.seed.json")

	if err != nil {
		panic(err)
	}

	err = ctx.Profiles.Save()

	if err != nil {
		panic(err)
	}

	err = ctx.Scopes.Seed("db/scopes.seed.json")

	if err != nil {
		panic(err)
	}

	err = ctx.Scopes.Save()

	if err != nil {
		panic(err)
	}
}
