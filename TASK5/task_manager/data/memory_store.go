package services
import 
	(
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
	return nil
} 

// Get the details of a specific task.
func (m *MemoryStore)Get(id string) (models.Task, error){
	return models.Task{}, nil
} 

// Get a list of all tasks.
func (m *MemoryStore)List() (map[string]models.Task, error){
	return m.tasks, nil
} 

// Update a specific task.
func (m *MemoryStore)Update(id string, task models.Task) (models.Task, error){
	return models.Task{}, nil
} 

//Delete a specific task.
func (m *MemoryStore)Remove(id string) error {
	return nil
}