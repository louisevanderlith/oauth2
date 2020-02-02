package core

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type SafeUser struct {
	Key         husk.Key
	Name        string
	Verified    bool
	DateCreated time.Time
	LastLogin   time.Time
}

func createSafeUser(user husk.Recorder) SafeUser {
	data := user.Data().(*User)
	meta := user.Meta()

	result := SafeUser{
		Key:         meta.Key,
		LastLogin:   data.LoginDate,
		Name:        data.Name,
		Verified:    data.Verified,
		DateCreated: time.Unix(0, meta.Key.Stamp),
	}

	return result
}

func GetUsers(page, size int) ([]SafeUser, error) {
	var result []SafeUser
	users, err := ctx.Users.Find(page, size, husk.Everything())

	if err != nil {
		return result, err
	}

	itor := users.GetEnumerator()

	for itor.MoveNext() {
		currUser := itor.Current()

		sfeUser := createSafeUser(currUser)
		result = append(result, sfeUser)
	}

	return result, nil
}
