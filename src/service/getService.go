package service

import "github.com/PasswordManager/remotedbadapter"

//GetPassword ...
//Handles get password to request it to the mongoadapter
func GetPassword(service string) (string, error) {
	pwd, err := remotedbadapter.SearchPassword(service, conf)
	if err != nil {
		return "", err
	}
	return pwd, nil
}
