package service

import "github.com/PasswordManager/remotedbadapter"

//GetItem ...
//Handles get item to request it to the mongoadapter
func GetItem(service string, itemdesc string) (string, error) {
	serviceuuid, err := SearchUUID(service)
	if err != nil {
		return "", err
	}

	item, err := remotedbadapter.SearchItem(serviceuuid, itemdesc, conf)
	if err != nil {
		return "", err
	}

	plainitem, err := decrypt(item)
	if err != nil {
		return "", err
	}

	return plainitem, nil
}
