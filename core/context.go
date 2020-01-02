package core

import "github.com/louisevanderlith/husk"

type context struct {
	Clients husk.Tabler
	Users husk.Tabler
	Forgotten husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Clients: husk.NewTable(Client{}),
		Users: husk.NewTable(User{}),
		Forgotten: husk.NewTable(Forgot{}),
	}

	seed()
}

func Shutdown() {
	ctx.Users.Save()
	ctx.Clients.Save()
}

func seed() {
	err := ctx.Users.Seed("db/users.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Users.Save()

	err = ctx.Clients.Seed("db/clients.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Clients.Save()
}