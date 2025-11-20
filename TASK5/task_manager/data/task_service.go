package services

import (
	"errors"

	"github.com/KimFarida/task_manager/models"
	"github.com/google/uuid"
)


type TaskService struct{
	// implements the interface TaskManager
	Store models.TaskManager 
}

func NewTaskService(store models.TaskManager) *TaskService{
	return &TaskService{Store: store}
}


// Here we can call all things that our Taskmanager Interface Implements
func (s *TaskService) CreateTask(t models.Task) error {
	if t.Title == ""{
		return errors.New("title is required")
	}

	if t.Description == ""{
		return errors.New("a short description is required")
	}

	if !models.IsValidStatus(t.Status){
		t.Status = models.TaskPending
	}
	
	t.ID = uuid.New().String()
	
	return s.Store.Add(t)
}

func (s *TaskService) GetTaskById(id string)(models.Task, error) {
	return s.Store.Get(id)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.Store.List()
}

func (s *TaskService) UpdateTask(id string, t models.Task) error {
	return s.Store.Update(id, t)
}


func (s *TaskService) DeleteTask(id string) error {
	return s.Store.Remove(id)
}
