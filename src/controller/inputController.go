package controller

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PasswordManager/configuration"
	"golang.org/x/crypto/bcrypt"
)

//Input ...
//User input structure
type Input struct {
	Command    string
	Subcommand string
	Arguments  []string
}

//HandleUserInput ...
//Handles User input for command execution
func (input *Input) HandleUserInput() {
	switch command := os.Args[1]; command {
	case "init":
		handleInitCommand()
	case "get":
		// handleGetCommand(os.Args[2])
	case "set":
		// handleSetCommand(os.Args[2])
	case "edit":
		// handleEditCommand(os.Args[2])
	case "import":
		// handleimportCommand(os.Args[2])
	case "export":
		// handleExportCommand(os.Args[2])
	default:
		return
	}
}

func handleInitCommand() (bool, string) {
	conf := new(configuration.Configuration)
	err := conf.LoadConfiguration()
	if err != nil {
		return false, err.Error()
	}

	if conf.Password != "" {
		return false, "Password already configured. No need to init again"
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter password: ")
	bPassword, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error()
	}

	fmt.Print("Confirm password: ")
	b2Password, err := reader.ReadString('\n')
	if err != nil {
		return false, err.Error()
	}

	if bPassword != b2Password {
		return false, "Passwords doesn't match!"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(bPassword), conf.Cost)
	if err != nil {
		return false, err.Error()
	}
	fmt.Println(string(hash))

	return true, ""
}

// func handleGetCommand(subcommand string) (bool, string) {

// }

// func handleSetCommand(subcommand string) (bool, string) {

// }

// func handleEditCommand(subcommand string) (bool, string) {

// }

// func handleimportCommand(subcommand string) (bool, string) {

// }

// func handleExportCommand(subcommand string) (bool, string) {

// }
