package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/PasswordManager/errorsdef"
	"github.com/PasswordManager/service"
	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
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
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%] %time% | %msg%",
		},
	}

	switch command := os.Args[1]; command {
	case "init":
		if exec, why, lvl := handleInitCommand(); !exec {
			if lvl == "ERR" {
				log.Error(why)
			} else {
				log.Warning(why)
			}
		} else {
			log.Info("PM initialization successful")
		}
	case "get":
		if len(os.Args) < 3 {
			missingSubCommand()
			return
		}
		if exec, why, lvl := handleGetCommand(os.Args[2]); !exec {
			if lvl == "ERR" {
				log.Error(why)
			} else {
				log.Warning(why)
			}
		} else {
			log.Info(why)
		}
	case "set":
		if len(os.Args) < 3 {
			missingSubCommand()
			return
		}
		if exec, why, lvl := handleSetCommand(os.Args[2]); !exec {
			if lvl == "ERR" {
				log.Error(why)
			} else {
				log.Warning(why)
			}
		} else {
			log.Info(why)
		}
	case "import":
		// handleimportCommand(os.Args[2])
	case "export":
		// handleExportCommand(os.Args[2])
	case "update":
		if len(os.Args) < 3 {
			missingSubCommand()
			return
		}
		if exec, why, lvl := handleUpdateCommand(os.Args[2]); !exec {
			if lvl == "ERR" {
				log.Error(why)
			} else {
				log.Warning(why)
			}
		} else {
			log.Info(why)
		}
	case "delete":
		if len(os.Args) < 3 {
			missingSubCommand()
			return
		}
		if exec, why, lvl := handleDeleteCommand(os.Args[2]); !exec {
			if lvl == "ERR" {
				log.Error(why)
			} else {
				log.Warning(why)
			}
		} else {
			log.Info(why)
		}
	case "help":
		printUsage()
	case "generate":
	case "list":
		if len(os.Args) < 1 {
			missingSubCommand()
			return
		}
		if exec, why, lvl := handleListCommand(); !exec {
			if lvl == "ERR" {
				log.Error(why)
			} else {
				log.Warning(why)
			}
		}
	default:
		commandNotFound()
		return
	}
}

func handleInitCommand() (bool, string, string) {
	return service.InitService()
}

func handleListCommand() (bool, string, string) {
	err := checkMasterPassword()
	if err != nil {
		return false, errorsdef.Mpassincorrect, "ERR"
	}

	fmt.Println("Services available")
	fmt.Println("---------------------------")
	for _, item := range *service.ListServices() {
		fmt.Println(item)
	}
	fmt.Println("")

	return true, "", ""
}

func handleGetCommand(subcommand string) (bool, string, string) {
	if len(os.Args) < 3 {
		printUsageSubCommands()
		return false, errorsdef.Missingparams, "ERR"
	}

	if subcommand == "password" {
		password := os.Args[3]
		err := checkMasterPassword()
		if err != nil {
			return false, errorsdef.Mpassincorrect, "ERR"
		}

		pwd, err := service.GetItem(password, subcommand)
		if err != nil {
			return false, err.Error(), "ERR"
		}
		clipboard.WriteAll(pwd)
		return true, "password copied to the clipboard!", "INFO"
	} else if subcommand == "username" {
		username := os.Args[3]
		err := checkMasterPassword()
		if err != nil {
			return false, errorsdef.Mpassincorrect, "ERR"
		}

		user, err := service.GetItem(username, subcommand)
		if err != nil {
			return false, err.Error(), "ERR"
		}
		clipboard.WriteAll(user)
		return true, "user copied to the clipboard!", "INFO"
	} else {
		printUsageSubCommands()
		return false, errorsdef.Missingparams, "ERR"
	}
}

func handleSetCommand(subcommand string) (bool, string, string) {
	if len(os.Args) < 6 {
		printUsageSubCommands()
		return false, errorsdef.Missingparams, "ERR"
	}

	serv, username, password := os.Args[3], os.Args[4], os.Args[5]
	// err := checkMasterPassword()
	// if err != nil {
	// 	return false, errorsdef.Mpassincorrect, "ERR"
	// }

	result, err := service.SetService(serv, username, password)
	if err != nil {
		return false, err.Error(), "ERR"
	}
	return true, result, "INFO"
}

// func handleimportCommand(subcommand string) (bool, string) {

// }

// func handleExportCommand(subcommand string) (bool, string) {

// }

func handleUpdateCommand(subcommand string) (bool, string, string) {
	if len(os.Args) < 3 {
		printUsageSubCommands()
		return false, errorsdef.Missingparams, "ERR"
	}

	if subcommand == "password" {
		serv := os.Args[3]
		err := checkMasterPassword()
		if err != nil {
			return false, errorsdef.Mpassincorrect, "ERR"
		}

		newpwd, err := askNewPassword()
		if err != nil {
			return false, err.Error(), "ERR"
		}

		result, err := service.UpdateItem(serv, newpwd, subcommand)
		if err != nil {
			return false, err.Error(), "ERR"
		}
		return true, result, "INFO"

	} else if subcommand == "username" {
		serv := os.Args[3]
		err := checkMasterPassword()
		if err != nil {
			return false, errorsdef.Mpassincorrect, "ERR"
		}

		newuser, err := askForInput("new username")
		if err != nil {
			return false, err.Error(), "ERR"
		}

		result, err := service.UpdateItem(serv, newuser, subcommand)
		if err != nil {
			return false, err.Error(), "ERR"
		}
		return true, result, "INFO"
	} else {
		printUsageSubCommands()
		return false, errorsdef.Subcommandnotfound, "ERR"
	}
}

func handleDeleteCommand(subcommand string) (bool, string, string) {
	if len(os.Args) < 3 {
		printUsageSubCommands()
		return false, errorsdef.Missingparams, "ERR"
	}

	serv := os.Args[3]
	err := checkMasterPassword()
	if err != nil {
		return false, errorsdef.Mpassincorrect, "ERR"
	}

	confirmation := askForConfirmation()
	if confirmation == "Y" || confirmation == "y" {
		result, err := service.DeleteService(serv)
		if err != nil {
			return false, err.Error(), "ERR"
		}
		return true, result, "INFO"
	}

	return true, "Operation canceled!", "INFO"
}

func checkMasterPassword() error {
	fmt.Print("Enter master password: ")
	bpass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	fmt.Println("")

	return service.CheckPasswords(string(bpass))
}

func askForInput(input string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(fmt.Sprintf("Enter %s: ", input))
	userinput, err := reader.ReadString('\n')
	service.CheckInput(&userinput)
	if err != nil {
		return "", err
	}
	return userinput, nil
}

func askNewPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	fPassword, err := reader.ReadString('\n')
	service.CheckInput(&fPassword)
	if err != nil {
		return "", err
	}

	fmt.Print("Confirm password: ")
	sPassword, err := reader.ReadString('\n')
	service.CheckInput(&sPassword)
	if err != nil {
		return "", err
	}

	if fPassword != sPassword {
		return "", errors.New("passwords doesn't match")
	}
	return fPassword, nil
}

func askForConfirmation() string {
	answer := ""
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Are you really sure? (Y/n): ")
	answer, err := reader.ReadString('\n')
	service.CheckInput(&answer)
	if err != nil {
		return ""
	}

	return answer
}

func commandNotFound() {
	fmt.Printf("Command not recognized %s - Usage of pm:\n\n", os.Args[1])
	printUsage()
}

func subCommandNotFound() {
	fmt.Printf("SubCommand not recognized %s - Usage of pm:\n\n", os.Args[2])
	printUsageSubCommands()
}

func missingSubCommand() {
	fmt.Printf("SubCommand is missing for this command %s - Usage of pm:\n\n", os.Args[1])
	printUsageSubCommands()
}

func printUsage() {
	fmt.Printf("\t$> pm COMMAND [SUBCOMMAND] search\n\n")
	fmt.Printf("Commands available:\n\texport\n\tget\n\thelp\n\timport\n\tinit\n\tset\n\tupdate\n\n")
	fmt.Printf("Example: $> pm get password gmail")
}

func printUsageSubCommands() {
	switch os.Args[1] {
	case "get":
		fmt.Printf("get | Return wanted information saved previously\n")
		fmt.Printf("\t$> pm get [SUBCOMMAND] search\n")
		fmt.Printf("Subcommands available:\n\tpassword\n\tusername\n\n")
		fmt.Printf("Example: $> pm get password gmail\n\n")
	case "update":
		fmt.Printf("update | update information\n")
		fmt.Printf("\t$> pm update [SUBCOMMAND] search\n")
		fmt.Printf("Subcommands available:\n\tpassword\n\tusername\n\n")
		fmt.Printf("Example: $> pm update username gmail\n\n")
	case "set":
		fmt.Printf("set | set a new service\n")
		fmt.Printf("\t$> pm set [SUBCOMMAND] servicename username password\n")
		fmt.Printf("Subcommands available:\n\tservice\n\n")
		fmt.Printf("Example: $> pm set service gmail example@yahoo.es pass1234\n\n")
	}
}
