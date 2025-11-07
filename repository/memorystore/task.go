package memorystore

import "todo_cli/entity"

type Task struct {
	tasks []entity.Task
}

func NewTaskStore() *Task {
	return &Task{
		tasks: make([]entity.Task, 0),
	}
}
func (t *Task) CreateNewTask(task entity.Task) (entity.Task, error) {
	task.Id = len(t.tasks) + 1
	t.tasks = append(t.tasks, task)
	return task, nil
}

func (t *Task) ListTasks(userId int) ([]entity.Task, error) {
	var userTasks []entity.Task
	for _, task := range t.tasks {
		if task.UserId == userId {
			userTasks = append(userTasks, task)
		}
	}
	return userTasks, nil
}

func (t *Task) DoesThisUserHaveThisCategoryID(userId, categoryId int) bool {
	for _, task := range t.tasks {
		if task.UserId == userId && task.CategoryId == categoryId {
			return true
		}
	}
	return false
}
