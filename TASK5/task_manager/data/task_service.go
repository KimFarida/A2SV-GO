package services

import 
(
	"github.com/KimFarida/task_manager/models"
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
	return nil
}