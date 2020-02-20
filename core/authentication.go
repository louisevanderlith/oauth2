package core

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

// password hashing cost
const cost int = 11

// Login will attempt to authenticate a user
func Login(username, password string) (string, error) {
	log.Printf("Logging in %s\r\n", username)
	//ip := authReq.App.IP
	//location := authReq.App.Location

	if len(password) < 6 {
		return "", errors.New("authentication failed")
	}

	if !strings.Contains(username, "@") {
		return "", errors.New("authentication failed")
	}

	userRec, err := authStore.Users.FindFirst(byEmail(username))

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
	err = authStore.Users.Update(userRec)

	if err != nil {
		return "", err
	}

	err = authStore.Users.Save()

	if err != nil {
		return "", err
	}

	if !passed {
		return "", errors.New("authentication failed")
	}

	return userRec.GetKey().String(), nil
}
