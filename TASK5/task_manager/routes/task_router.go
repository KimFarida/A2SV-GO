package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/KimFarida/task_manager/controllers"
	"net/http"


)


func RegisterRoutes(r *gin.Engine, taskController *controllers.TaskController){
	r.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message":"WELCOME TO TASK MANAGER"})
	})

	// Grouping Routes 
	task := r.Group("tasks")
	{
		//	GET /tasks: Get a list of all tasks.
		task.GET("/", taskController.GetAllTasks)

		// GET /tasks/:id: Get the details of a specific task.
		task.GET("/:id", taskController.GetTaskById)

		// POST /tasks: Create a new task,
		// accept a JSON body with the task's title,
		// description, due date, and status.
		task.POST("/", taskController.CreateTask)
		
		// PUT /tasks/:id: Update a specific task.
		// accept a JSON body with the new details of the task.
		task.PUT("/:id", taskController.UpdateTask)

		// DELETE /tasks/:id: Delete a specific task.
		task.DELETE("/:id", taskController.DeleteTask)
	}
}