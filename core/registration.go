package core

import (
	"errors"

	"github.com/louisevanderlith/husk"
)

type Registration struct {
	ClientID       string
	Name           string
	Email          string
	Password       string
	PasswordRepeat string
}

func Register(r Registration) (husk.Recorder, error) {
	if r.Password != r.PasswordRepeat {
		return nil, errors.New("passwords do not match")
	}

	if len(r.ClientID) == 0 {
		return nil, errors.New("client id can not be empty")
	}

	if emailExists(r.Email) {
		return nil, errors.New("email already in use")
	}

	user, err := NewUser(r.Name, r.Email)

	if err != nil {
		return nil, err
	}

	user.SecurePassword(r.Password)
	user.AddTrace(getRegistrationTrace(r))

	//Expand registration to add Permissions for API also. Won't always be a 'User'
	//user.AddRole(r.App.Name, roletype.User)

	rec := ctx.Users.Create(user)
	defer ctx.Users.Save()

	if rec.Error != nil {
		return nil, rec.Error
	}

	return rec.Record, nil
}
