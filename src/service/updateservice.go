package service

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/PasswordManager/remotedbadapter"
)

//UpdatePassword ...
//Update a password of a existing service
func UpdatePassword(service string) (string, error) {
	serviceuuid, err := SearchUUID(service)
	if err != nil {
		return "", err
	}

	newpwd, err := askNewPassword()
	if err != nil {
		return "", err
	}

	cryptedpwd, err := encrypt(newpwd)
	if err != nil {
		return "", err
	}

	result, err := remotedbadapter.UpdatePassword(serviceuuid, cryptedpwd, conf)
	if err != nil {
		return "", err
	}
	return result, nil
}

func askNewPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	fPassword, err := reader.ReadString('\n')
	CheckInput(&fPassword)
	if err != nil {
		return "", err
	}

	fmt.Print("Confirm password: ")
	sPassword, err := reader.ReadString('\n')
	CheckInput(&sPassword)
	if err != nil {
		return "", err
	}

	if fPassword != sPassword {
		return "", errors.New("passwords doesn't match")
	}
	return fPassword, nil
}
