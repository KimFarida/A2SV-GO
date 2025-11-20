package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/KimFarida/task_manager/data"
)

type TaskController struct {
	Service *services.TaskService
}

// Initialize a new Task Controller
func NewTaskController(s *services.TaskService) *TaskController{
	return &TaskController{Service: s}
}




func (c *TaskController) CreateTask(ctx *gin.Context){}

func (c *TaskController) UpdateTask(ctx *gin.Context){}

func (c *TaskController) GetTaskById(ctx *gin.Context){}

func (c *TaskController) GetAllTasks(ctx *gin.Context){}

func (c *TaskController) DeleteTask(ctx *gin.Context){}