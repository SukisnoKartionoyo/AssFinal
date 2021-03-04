package main

import (
	"assfinal/config"
	"assfinal/controllers"

	_ "assfinal/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Ass Final API
// @version 1.0
// @description This is a basic API Todo of which has proprerties of Project, Label, Status, and Title using Gin and Gorm.

// @contact.name KERJAAN FINAL
// @contact.email sukisno.kartionoyo@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()
	config.ConnectDB()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/task", controllers.GetAllTask)
		v1.GET("/task/:task_id", controllers.GetTaskByID)
		v1.PUT("/task/:task_id", controllers.UpdateTaskByID)
		v1.DELETE("/task/:task_id", controllers.DeleteTaskByID)
		v1.POST("/task", controllers.CreateNewTask)
		v1.POST("/project", controllers.CreateProject)
		v1.POST("/label", controllers.CreateLabel)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
