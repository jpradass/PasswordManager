package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	guuid "github.com/google/uuid"
)

const emptyUUID = "00000000-0000-0000-0000-000000000000"

//UUIDService ...
//Represents services.json
type uuidservice struct {
	CreatedAt time.Time  `json:"createdAt"`
	UUID      guuid.UUID `json:"uuid"`
}

var (
	services map[string]uuidservice = make(map[string]uuidservice)
	expath   string
)

func init() {
	ex, _ := os.Executable()
	expath = filepath.Dir(ex)

	byteValue, _ := ioutil.ReadFile(fmt.Sprintf("%s/configuration/services.json", expath))
	json.Unmarshal(byteValue, &services)
}

//SearchUUID ...
//Finds a coincidence service
func SearchUUID(service string) (string, error) {
	s, ok := services[service]
	if !ok {
		return "", errors.New("No service found with this name: " + service)
	}
	return s.UUID.String(), nil
}

//ExistsService ...
//Search if service already exists
func ExistsService(service string) error {
	_, ok := services[service]
	if ok {
		return errors.New("There is already a service with the name: " + service)
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
	services[service] = uuidservice{
		time.Now(),
		suuid,
	}

	err = saveServicesJSON()
	if err != nil {
		return "", err
	}

	return suuid.String(), nil
}

//DelService ...
//Delete given service
func DelService(service string) (string, error) {
	suuid := services[service].UUID.String()
	delete(services, service)

	err := saveServicesJSON()
	if err != nil {
		return "", err
	}

	return suuid, nil
}

func saveServicesJSON() error {
	out, err := json.Marshal(services)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/configuration/services.json", expath), out, 0644)
	if err != nil {
		return err
	}

	return nil
}
