package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"todo_cli/contract"
	"todo_cli/repository/memorystore"
	"todo_cli/service/task"

	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo_cli/entity"
	"todo_cli/filestore"
)

func inputScanner(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(text)
	scanner.Scan()
	return scanner.Text()
}

var userStorge = []entity.User{}
var TaskStorge = []entity.Task{}
var AthenticatedUser *entity.User
var CategoryStorge = []entity.Category{}

func main() {
	taskMemoryRepo := memorystore.NewTaskStore()
	taskService := task.NewService(taskMemoryRepo)
	command := flag.String("command", "no command", "create a new task")
	flag.Parse()
	// scanner := bufio.NewScanner(os.Stdin)
	for {
		runCommand(*command, &taskService)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command:")
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string, taskService *task.Service) {
	if command != "register" && command != "exit" && AthenticatedUser == nil {
		login()
		if AthenticatedUser == nil {
			return
		}
	}
	var userFileStore = filestore.New("./user.txt")
	userLoadFromStorage(userFileStore)
	var store contract.UserWriteStore

	store = userFileStore
	switch command {
	case "create-task":
		createTask(taskService)
	case "create-category":
		createCategory()
	case "task-list":
		taskList(taskService)
	case "category-list":
		categoryList()
	case "register":
		register(store)
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func register(store contract.UserWriteStore) {
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

	user := entity.User{
		Id:       len(userStorge) + 1,
		Name:     name,
		Password: passwordHasher(password),
		Email:    email,
	}
	userStorge = append(userStorge, user)

	store.Save(user)
	login()

}
func createCategory() {
	title := inputScanner("set a title for category:")
	color := inputScanner("set a color for category:")
	newcat := entity.Category{Id: len(CategoryStorge) + 1,
		Title:  title,
		Color:  color,
		UserId: AthenticatedUser.Id}
	CategoryStorge = append(CategoryStorge, newcat)
}
func createTask(taskService *task.Service) {

	title := inputScanner("set a title:")
	category := inputScanner("set a category id:")
	category = strings.TrimSpace(category)
	categoryId, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println(err)
		return
	}
	iffound := false
	for _, c := range CategoryStorge {
		if c.Id == categoryId && c.UserId == AthenticatedUser.Id {
			iffound = true
			break

		}
	}
	if !iffound {
		fmt.Println("Category not found")
		return
	}
	//task := entity.Task{
	//	Id:         len(TaskStorge) + 1,
	//	Title:      title,
	//	DueDate:    "",
	//	CategoryId: 0,
	//	IsDone:     false,
	//	UserId:     AthenticatedUser.Id,
	//}

	//TaskStorge = append(TaskStorge, task)
	response, err := taskService.CreateTask(task.CreateRequest{
		Title:               title,
		DueDate:             "",
		CategoryId:          categoryId,
		AuthenticatedUserId: AthenticatedUser.Id,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)

}
func taskList(taskService *task.Service) {
	userTasks, err := taskService.List(AthenticatedUser.Id)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Println("user tasks:", userTasks)
}
func categoryList() {
	for _, cate := range CategoryStorge {
		fmt.Printf("%d %s\n", cate.Id, cate.Title)
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
			if passwordHasher(password) == user.Password {
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

func userLoadFromStorage(store contract.UserReadStore) {
	users := store.Load()
	userStorge = append(userStorge, users...)

}

func passwordHasher(password string) string {
	hash := md5.Sum([]byte(password))

	// تبدیل هش به رشته‌ی hex قابل چاپ
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
