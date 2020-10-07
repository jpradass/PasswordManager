package controller

import (
	"fmt"
	"os"

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
	case "edit":
		// handleEditCommand(os.Args[2])
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
	case "help":
		printUsage()
	default:
		commandNotFound()
		return
	}
}

func handleInitCommand() (bool, string, string) {
	return service.InitService()
}

func handleGetCommand(subcommand string) (bool, string, string) {
	if subcommand == "password" {
		pwd, err := service.GetPassword(os.Args[3])
		if err != nil {
			return false, err.Error(), "ERR"
		}
		clipboard.WriteAll(pwd)
		return true, "password copy to the clipboard!", "INFO"
	} else if subcommand == "username" {

	}

	return true, "", ""
}

func handleSetCommand(subcommand string) (bool, string, string) {
	result, err := service.SetService("gmail", "pepes", "hola")
	if err != nil {
		return false, err.Error(), "ERR"
	}
	return true, result, "INFO"
}

// func handleEditCommand(subcommand string) (bool, string) {

// }

// func handleimportCommand(subcommand string) (bool, string) {

// }

// func handleExportCommand(subcommand string) (bool, string) {

// }

func handleUpdateCommand(subcommand string) (bool, string, string) {
	if subcommand == "password" {
		fPassword := "hola"
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Enter password: ")
		// fPassword, err := reader.ReadString('\n')
		// if err != nil {
		// 	return false, err.Error(), "ERR"
		// }

		// fmt.Print("Confirm password: ")
		// sPassword, err := reader.ReadString('\n')
		// if err != nil {
		// 	return false, err.Error(), "ERR"
		// }

		// if fPassword != sPassword {
		// 	return false, "Passwords doesn't match!", "WARN"
		// }
		service.UpdatePassword(os.Args[3], fPassword)
	} else if subcommand == "username" {

	}

	return true, "", ""
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
		fmt.Printf("\t$> pm get [SUBCOMMAND] search\n\n")
		fmt.Printf("Subcommands available:\n\tpassword\n\tusername\n\n")
		fmt.Printf("Example: $> pm get password gmail")
	case "update":
		fmt.Printf("update | update information\n")
		fmt.Printf("\t$> pm update [SUBCOMMAND] search\n\n")
		fmt.Printf("Subcommands available:\n\tpassword\n\tusername\n\n")
		fmt.Printf("Example: $> pm update username gmail")
	}
}
