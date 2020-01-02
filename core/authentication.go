package core

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// password hashing cost
const cost int = 11

// Login will attempt to authenticate a user
func Login(username, password string) (string, error) {
	//ip := authReq.App.IP
	//location := authReq.App.Location

	if len(password) < 6 {
		return "", errors.New("authentication failed")
	}

	if !strings.Contains(username, "@") {
		return "", errors.New("authentication failed")
	}

	userRec, err := getUser(username)

	if err != nil {
		return "", err
	}

	user := userRec.Data().(User)

	if !user.Verified {
		return "", errors.New("authentication failed")
	}

	compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	passed := compare == nil
	user.AddTrace(getLoginTrace(passed))
	err = ctx.Users.Update(userRec)

	if err != nil {
		return "", err
	}

	err = ctx.Users.Save()

	if err != nil {
		return "", err
	}

	if !passed {
		return "", errors.New("authentication failed")
	}

	return userRec.GetKey().String(), nil
}
