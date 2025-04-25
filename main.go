package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	Name     string
	Id       int
	Email    string
	Password string
}

var userStorge = []User{}
var AthenticatedUser *User

func main() {

	command := flag.String("command", "no command", "create a new task")
	flag.Parse()
	// scanner := bufio.NewScanner(os.Stdin)
	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command:")
		scanner.Scan()
		*command = scanner.Text()
	}

}
func runCommand(command string) {
	if command != "register" && command != "exit" && AthenticatedUser == nil {
		login()
	}
	switch command {
	case "create-task":
		createTask()
	case "register":
		register()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}
func register() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Register")
	fmt.Println("Enter your email:")
	scanner.Scan()
	email := scanner.Text()

	fmt.Println("Enter your name:")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter your password:")
	scanner.Scan()
	password := scanner.Text()
	user := User{
		Id:       len(userStorge) + 1,
		Name:     name,
		Password: password,
		Email:    email,
	}
	userStorge = append(userStorge, user)
	login()

}

func createTask() {

}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Login")
	fmt.Println("Enter your email:")
	scanner.Scan()
	email := scanner.Text()

	fmt.Println("Enter your password:")
	scanner.Scan()
	password := scanner.Text()

	notfound := true
	for _, user := range userStorge {
		if email == user.Email {
			if password == user.Password {
				notfound = false
				AthenticatedUser = &user
				fmt.Println("you are login!!")

				break

			} else {
				fmt.Println("your password is not correct.")
			}
		}

	}
	if notfound {
		fmt.Print("cant find user or password")
		return
	}
}

// fmt.Println("hello")
// 	command := flag.String("command", "no command", "create a new task")
// 	flag.Parse()
// 	fmt.Println(*command)
// 	scanner := bufio.NewScanner(os.Stdin)
// 	if *command == "create-task" {
// 		fmt.Println(":")

// 		scanner.Scan()
// 		name := scanner.Text()
// 		fmt.Println(name)
// 	}
