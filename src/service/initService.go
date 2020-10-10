package service

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

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

	_, err := os.Stat("configuration/services.json")
	if err != nil {
		f, err := os.OpenFile("configuration/services.json", os.O_CREATE, 0644)
		if err != nil {
			return false, err.Error(), "ERR"
		}
		f.Close()
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter database hostname/ip address: ")
	host, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	checkInput(&host)
	conf.Host = host

	fmt.Print("Enter database username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	checkInput(&username)
	conf.User = username

	fmt.Print("Enter database password: ")
	dbpassword, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	checkInput(&dbpassword)
	conf.DbPass = dbpassword

	fmt.Print("Enter database name: ")
	db, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	checkInput(&db)
	conf.DB = db

	fmt.Print("Enter collection name: ")
	collection, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	checkInput(&collection)
	conf.Collection = collection

	fmt.Print("Enter master password: ")
	bPassword, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	checkInput(&bPassword)

	fmt.Print("Confirm master password: ")
	b2Password, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	checkInput(&b2Password)

	if bPassword != b2Password {
		return false, "Passwords doesn't match!", "WARN"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(bPassword), conf.Cost)
	if err != nil {
		return false, err.Error(), "ERR"
	}
	conf.Password = string(hash)
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return false, err.Error(), "ERR"
	}
	conf.Key = base64.StdEncoding.EncodeToString(key)[:32]

	fmt.Print("Enter bcrypt algorithm cost (by default 12): ")
	cost, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error(), "ERR"
	}
	checkInput(&cost)
	if len(cost) == 0 {
		conf.Cost = 12
	} else {
		conf.Cost, _ = strconv.Atoi(cost)
	}

	err = conf.SaveConfiguration()
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

func checkInput(input *string) {
	if runtime.GOOS == "windows" {
		*input = strings.Replace(*input, "\r\n", "", -1)
	}
}
