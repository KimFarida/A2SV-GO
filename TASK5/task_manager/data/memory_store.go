package services

import (
	"errors"
	"fmt"

	"github.com/KimFarida/task_manager/models"
)


type MemoryStore struct {
	tasks map[string]models.Task
}

// Constructor returns a new instance of memory 
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: make(map[string]models.Task),
	}
}

// now implement the interface 
// Create a new task.
func (m *MemoryStore)Add(task models.Task) error{
	m.tasks[task.ID] = task
	return nil
} 

// Get the details of a specific task.
func (m *MemoryStore)Get(id string) (models.Task, error){
	if task, ok := m.tasks[id]; ok{
		return task, nil
	}
	return models.Task{}, errors.New("this task does not exist")
} 

// Get a list of all tasks.
func (m *MemoryStore)List() ([]models.Task, error){
	if m.tasks == nil{
		return make([]models.Task, 0), nil
	}

	allTasks := []models.Task{}
	for _, task := range m.tasks{
		allTasks = append(allTasks, task)
	}

	return allTasks, nil
}

// Update a specific task.
func (m *MemoryStore)Update(id string, task models.Task) error{
	//taskToUpdate
	existingTask, ok :=  m.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	
	if task.Title != ""{
		existingTask.Title = task.Title
	} 
	if task.Description != ""{
		existingTask.Description = task.Description
	}

	if models.IsValidStatus(task.Status){
		existingTask.Status = task.Status
	}
	
	if !task.DueDate.IsZero(){
		existingTask.DueDate = task.DueDate
	}
	m.tasks[id] = existingTask
	return nil
} 

//Delete a specific task.
func (m *MemoryStore)Remove(id string) error {
	if _, ok := m.tasks[id]; !ok {
		return fmt.Errorf("task not found")
	}

	delete(m.tasks, id)
	return nil
}