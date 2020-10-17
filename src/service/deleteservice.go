package service

import (
	"fmt"

	"github.com/PasswordManager/remotedbadapter"
)

//DeleteService ...
//Deletes a given service
func DeleteService(service string) (string, error) {
	err := ExistsService(service)
	if err == nil {
		return "", fmt.Errorf("No service found with this name: %s", service)
	}

	suuid, err := DelService(service)
	if err != nil {
		return "", err
	}

	result, err := remotedbadapter.RemoveService(suuid, conf)
	if err != nil {
		return "", err
	}

	return result, nil
}
