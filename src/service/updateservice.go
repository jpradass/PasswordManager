package service

import "github.com/PasswordManager/remotedbadapter"

//UpdatePassword ...
//Update a password of a existing service
func UpdatePassword(service string, pwd string) (string, error) {
	cryptedpwd, err := encrypt(pwd)
	if err != nil {
		return "", err
	}

	result, err := remotedbadapter.UpdatePassword([]byte(service), cryptedpwd, conf)
	if err != nil {
		return "", err
	}
	return result, nil
}
