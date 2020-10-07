package service

import (
	"github.com/PasswordManager/remotedbadapter"
)

//SetService ...
//Set a new service to the db
func SetService(service string, username string, pwd string) (string, error) {
	// TODO Hash the service so its not so visible as base64
	cryptedpwd, err := encrypt(pwd)
	if err != nil {
		return "", err
	}

	crypteduser, err := encrypt(username)
	if err != nil {
		return "", err
	}

	result, err := remotedbadapter.InsertService([]byte(service), crypteduser, cryptedpwd, conf)
	if err != nil {
		return "", err
	}
	return result, nil
}
