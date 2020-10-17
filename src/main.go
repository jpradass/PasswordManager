package main

import "github.com/PasswordManager/controller"

func main() {
	controller := new(controller.Input)
	controller.HandleUserInput()
}
