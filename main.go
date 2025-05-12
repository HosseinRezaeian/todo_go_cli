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
type Task struct {
	Id         int
	Title      string
	DueDate    string
	CategoryId int
	IsDone     bool
	UserId     int
}
type Category struct {
	Id     int
	Name   string
	Color  string
	UserId int
}

func (u User) print() {
	fmt.Println("Name:", u.Name, "Id:", u.Id, "Email:", u.Email)
}
func inputScanner(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(text)
	scanner.Scan()
	return scanner.Text()
}

var userStorge = []User{}
var TaskStorge = []Task{}
var AthenticatedUser *User
var CategoryStorge = []Category{}

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
		if AthenticatedUser == nil {
			return
		}
	}
	switch command {
	case "create-task":
		createTask()
	case "task-list":
		taskList()
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
	if AthenticatedUser != nil {
		AthenticatedUser.print()
	}
	title := inputScanner("set a title:")
	task := Task{
		Id:       len(TaskStorge) + 1,
		Title:    title,
		DueDate:  "",
		Category: "",
		IsDone:   false,
		UserId:   AthenticatedUser.Id,
	}
	TaskStorge = append(TaskStorge, task)
}
func taskList() {
	for _, task := range TaskStorge {
		fmt.Printf("%d %s\n", task.Id, task.Title)
	}
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
