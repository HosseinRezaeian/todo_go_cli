package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"todo_cli/entity"
)

type FileStore struct {
	filepath string
}

func New(path string) FileStore {
	return FileStore{filepath: path}
}
func (f FileStore) Save(user entity.User) {
	f.writeUser(user)
}

func (f FileStore) writeUser(user entity.User) {
	path := "user.txt"
	var file *os.File
	defer file.Close()
	_, err := os.Stat(path)

	fmt.Println(err)
	file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	newUser, err := json.Marshal(user)
	//newUser := fmt.Sprintf("id:%d ,name:%s ,email:%s,password:%s \n", user.Id, user.Name, user.Email, user.Password)
	var b = []byte(newUser)
	file.Write([]byte("\n"))

	file.Write(b)
}

func (f FileStore) Load() []entity.User {
	var uStorage []entity.User
	file, err := os.Open("user.txt")

	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	var data = make([]byte, 1024)
	numberOfBytes, err := file.Read(data)
	var datastr = string(data[:numberOfBytes])
	userSlice := strings.Split(datastr, "\n")
	fmt.Println("s", userSlice)
	//fmt.Println(userSlice)
	for _, u := range userSlice {
		user, err := deserializersUser(u)
		if err != nil {
			fmt.Println(err)
			continue
		}
		uStorage = append(uStorage, user)
	}

	return uStorage
}

func deserializersUser(userStr string) (entity.User, error) {
	if userStr == "" {
		return entity.User{}, errors.New("send user string")
	}
	userStuct := entity.User{}
	err := json.Unmarshal([]byte(userStr), &userStuct)
	return userStuct, err
}
