package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	guuid "github.com/google/uuid"
)

//ServiceUUID ...
//Represents services.json
type ServiceUUID struct {
	Service string     `json:"service"`
	UUID    guuid.UUID `json:"uuid"`
}

var services []ServiceUUID

func init() {
	byteValue, _ := ioutil.ReadFile("configuration/services.json")
	json.Unmarshal(byteValue, &services)
}

//SearchUUID ...
//Finds a coincidence service
func SearchUUID(service string) (string, error) {
	for _, s := range services {
		if strings.Contains(s.Service, service) {
			return s.UUID.String(), nil
		}
	}
	return "", errors.New("No service found with this name: " + service)
}

//ExistsService ...
//Search if service already exists
func ExistsService(service string) error {
	for _, s := range services {
		if s.Service == service {
			return errors.New("There is already a service with the name: " + service)
		}
	}
	return nil
}

//NewService ...
//Creates an uuid and associate it to the new service
func NewService(service string) (string, error) {
	err := ExistsService(service)
	if err != nil {
		return "", err
	}

	suuid := guuid.New()
	services = append(services, ServiceUUID{
		service,
		suuid,
	})

	out, err := json.Marshal(services)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile("configuration/services.json", out, 0644)
	if err != nil {
		return "", err
	}

	return suuid.String(), nil
}
