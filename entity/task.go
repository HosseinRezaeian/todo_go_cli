package entity

type Task struct {
	Id         int
	Title      string
	DueDate    string
	CategoryId int
	IsDone     bool
	UserId     int
}
