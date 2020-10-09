package service

import "github.com/PasswordManager/remotedbadapter"

//GetPassword ...
//Handles get password to request it to the mongoadapter
func GetPassword(service string) (string, error) {
	serviceuuid, err := SearchUUID(service)
	if err != nil {
		return "", err
	}

	pwd, err := remotedbadapter.SearchPassword(serviceuuid, conf)
	if err != nil {
		return "", err
	}

	plainpwd, err := decrypt(pwd)
	if err != nil {
		return "", err
	}

	return plainpwd, nil
}
