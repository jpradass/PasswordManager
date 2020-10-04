package service

import "github.com/PasswordManager/remotedbadapter"

//GetPassword ...
//Handles get password to request it to the mongoadapter
func GetPassword(service string) (string, error) {
	cryptedservice, err := encrypt(service)
	if err != nil {
		return "", err
	}

	pwd, err := remotedbadapter.SearchPassword(cryptedservice, conf)
	if err != nil {
		return "", err
	}

	pwd, err = decrypt(pwd)
	if err != nil {
		return "", err
	}

	return pwd, nil
}
