package models


type TaskManager interface{ 
	Add(task Task) error // Create a new task.
	Get(id string) (Task, error) // Get the details of a specific task.
	List() ([]Task, error) // Get a list of all tasks.
	Update(id string, task Task) error // Update a specific task. 
	Remove(id string) error //Delete a specific task.
}