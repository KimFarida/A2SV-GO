package controllers

import (
	//"errors"
	"net/http"

	"github.com/KimFarida/task_manager/data"
	"github.com/KimFarida/task_manager/models"
	"github.com/gin-gonic/gin"

)

type TaskController struct {
	Service *services.TaskService
}

// Initialize a new Task Controller
func NewTaskController(s *services.TaskService) *TaskController{
	return &TaskController{Service: s}
}



func (c *TaskController) CreateTask(ctx *gin.Context){
	var newTask models.Task
	if err:= ctx.BindJSON(&newTask) ; err != nil{
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid JSON"},
		)
		return
	}

	if err:= c.Service.CreateTask(newTask); err != nil{
		ctx.IndentedJSON(
			http.StatusBadRequest,
			 gin.H{"error": err.Error()},
			)
		return
	}
	
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message":"Task created successfuly"})
	

	
}
func (c *TaskController) GetTaskById(ctx *gin.Context){
	id := ctx.Param("id")

	if id == ""{
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Missing Parameter ID"},
		)
		return
	}

	task, err := c.Service.GetTaskById(id)

	if err != nil{
		ctx.IndentedJSON(
			http.StatusBadRequest,
			 gin.H{"error": err.Error()},
			)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"task":task})
}

func (c *TaskController) GetAllTasks(ctx *gin.Context){
	tasks, err := c.Service.GetTasks()

	if err != nil{
		ctx.IndentedJSON(
			http.StatusInternalServerError,
			 gin.H{"error": err.Error()},
			)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"tasks":tasks})
}

func (c *TaskController) UpdateTask(ctx *gin.Context){
	var task models.Task

	id := ctx.Param("id")

	if id == ""{
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Missing Parameter ID"},
		)
		return
	}

	if err:= ctx.BindJSON(&task) ; err != nil{
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid JSON"},
		)
		return
	}

	err := c.Service.UpdateTask(id, task)

	if err != nil{
		ctx.IndentedJSON(
			http.StatusNotFound,
			 gin.H{"error": err.Error()},
			)
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message":"Task Updated successfully"})
}


func (c *TaskController) DeleteTask(ctx *gin.Context){

	id := ctx.Param("id")

	if id == ""{
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Missing Parameter ID"},
		)
		return
	}


	err := c.Service.DeleteTask(id)

	if err != nil{
		ctx.IndentedJSON(
			http.StatusNotFound,
			 gin.H{"error": err.Error()},
			)
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message":"Task Deleted successfully"})
}
