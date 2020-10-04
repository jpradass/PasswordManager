package service

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PasswordManager/configuration"
	"golang.org/x/crypto/bcrypt"
)

var (
	conf *configuration.Configuration
	iv   = "1010101010101010"
)

func init() {
	conf = new(configuration.Configuration)
	conf.LoadConfiguration()
}

//InitService ...
//Init service setting up all data necessary
func InitService() (bool, string, string) {
	if conf.Password != "" {
		return false, "Data already configured. No need to init again", "WARN"
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter database hostname/ip address: ")
	host, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	conf.Host = host

	fmt.Print("Enter database username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	conf.User = username

	fmt.Print("Enter database password: ")
	dbpassword, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	conf.DbPass = dbpassword

	fmt.Print("Enter database name: ")
	db, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	conf.DB = db

	fmt.Print("Enter master password: ")
	bPassword, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}

	fmt.Print("Confirm master password: ")
	b2Password, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}

	if bPassword != b2Password {
		return false, "Passwords doesn't match!", "WARN"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(bPassword), conf.Cost)
	if err != nil {
		return false, err.Error(), "ERR"
	}

	err = conf.SetPassword(string(hash))
	if err != nil {
		return false, err.Error(), "ERR"
	}
	return true, "", ""
}

//CheckPasswords ...
//Check if master password match
func CheckPasswords(provided string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(conf.Password), []byte(provided)); err != nil {
		return false
	}
	return true
}
