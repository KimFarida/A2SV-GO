package main

import (
	"github.com/KimFarida/task_manager/controllers"
	"github.com/KimFarida/task_manager/data"
	"github.com/KimFarida/task_manager/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	router:= gin.Default()

	store := services.NewMemoryStore()
	taskService := services.NewTaskService(store)
	taskController := controllers.NewTaskController(taskService)

	routes.RegisterRoutes(router, taskController)

	router.Run("localhost:8080")
}