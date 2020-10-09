package service

import (
	"github.com/PasswordManager/remotedbadapter"
)

//SetService ...
//Set a new service to the db
func SetService(service string, username string, pwd string) (string, error) {
	serviceuuid, err := NewService(service)
	if err != nil {
		return "", err
	}

	cryptedpwd, err := encrypt(pwd)
	if err != nil {
		return "", err
	}

	crypteduser, err := encrypt(username)
	if err != nil {
		return "", err
	}

	result, err := remotedbadapter.InsertService(serviceuuid, crypteduser, cryptedpwd, conf)
	if err != nil {
		return "", err
	}
	return result, nil
}
