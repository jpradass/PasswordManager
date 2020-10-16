package service

import (
	"github.com/PasswordManager/remotedbadapter"
)

//UpdateItem ...
//Update a given item
func UpdateItem(service string, item string, itemdesc string) (string, error) {
	serviceuuid, err := SearchUUID(service)
	if err != nil {
		return "", err
	}

	crypteditem, err := encrypt(item)
	if err != nil {
		return "", err
	}

	result, err := remotedbadapter.UpdateItem(serviceuuid, crypteditem, itemdesc, conf)
	if err != nil {
		return "", err
	}
	return result, nil
}
