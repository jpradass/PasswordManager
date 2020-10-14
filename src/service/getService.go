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

//GetUsername ...
//Handles get user to request it to the mongoadapter
func GetUsername(service string) (string, error) {
	serviceuuid, err := SearchUUID(service)
	if err != nil {
		return "", err
	}

	username, err := remotedbadapter.SearchUsername(serviceuuid, conf)
	if err != nil {
		return "", err
	}

	plainuser, err := decrypt(username)
	if err != nil {
		return "", err
	}

	return plainuser, nil
}
