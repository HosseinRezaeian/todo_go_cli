package task

import (
	"errors"
	"fmt"
	"todo_cli/entity"
)

type ServiceRepository interface {
	DoesThisUserHaveThisCategoryID(userId, categoryId int) bool
	CreateNewTask(t entity.Task) (entity.Task, error)
	ListTasks(userId int) ([]entity.Task, error)
}
type Service struct {
	repository ServiceRepository
}

func NewService(repository ServiceRepository) Service {
	return Service{
		repository: repository,
	}
}

type CreateRequest struct {
	Title               string
	DueDate             string
	CategoryId          int
	AuthenticatedUserId int
}
type CreateResponse struct {
	Task entity.Task
}

func (t Service) CreateTask(req CreateRequest) (CreateResponse, error) {
	ok := t.repository.DoesThisUserHaveThisCategoryID(req.AuthenticatedUserId, req.CategoryId)
	if !ok {
		return CreateResponse{}, errors.New("User does not have this category")
	}

	task := entity.Task{
		Id:         0,
		Title:      req.Title,
		DueDate:    req.DueDate,
		CategoryId: req.CategoryId,
		UserId:     req.AuthenticatedUserId,
	}
	createdTask, err := t.repository.CreateNewTask(task)
	if err != nil {
		return CreateResponse{}, fmt.Errorf("Failed to create task: %v", err)

	}
	return CreateResponse{createdTask}, nil
}

type ListResponse struct {
	tasks []entity.Task
}

type ListRequest struct {
	UserId int
}

func (t Service) List(userId int) (ListResponse, error) {
	tasks, err := t.repository.ListTasks(userId)
	if err != nil {
		return ListResponse{}, fmt.Errorf("Failed to list tasks: %v", err)
	}
	return ListResponse{tasks}, nil
}
