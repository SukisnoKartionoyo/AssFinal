package controllers

import (
	"assfinal/config"
	"assfinal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTaskByID godoc
// @Summary Show task detail including project, label, status based on Task Id
// @Tags tasks
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param task_id path int true "Task ID"
// @success 200 {object} models.Task
// @Router /task/{task_id} [get]
func GetTaskByID(c *gin.Context) {
	taskID := c.Params.ByName("task_id")
	var task models.Task
	if err := config.DB.Preload(("Project")).Preload(("Label")).Preload(("Status")).Where("task_id", taskID).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// GetAllTask godoc
// @Summary Show a list of Tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @success 200 {object} models.Task
// @Router /task [get]
func GetAllTask(c *gin.Context) {
	var task []models.Task
	if err := config.DB.Preload("Project").Preload("Label").Preload("Status").Find(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateNewTask godoc
// @Summary Create task
// @Description Add by JSON Task
// @Tags task
// @Accept  json
// @Produce  json
// @Param task body models.Task true "Add Task"
// @success 200 {object} models.Task
// @Router /task [post]
func CreateNewTask(c *gin.Context) {

	var task models.Task
	c.ShouldBindJSON(&task)
	var project models.Project
	if err := config.DB.Find(&project).Where("project_title", task.ProjectID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project Cannot Be Found"})
		return
	}
	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Created"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTaskByID godoc
// @Summary Update task or items by TaskId
// @Description Update by JSON Task
// @Tags task
// @Accept  json
// @Produce  json
// @Param task_id path int true "Task Id"
// @Param task body models.Task true "Update Task"
// @success 200 {object} models.Task
// @Router /task_id/{task_id} [put]
func UpdateTaskByID(c *gin.Context) {
	taskID := c.Params.ByName("taskID")
	var task models.Task
	if err := config.DB.Preload(("Project")).Preload(("Label")).Preload(("Status")).Where("task_id", taskID).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.ShouldBindJSON(&task)
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Updated!"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// DeleteTaskByID godoc
// @Summary Delete task including project, label, status by task id
// @Tags task
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param task_id path int true "Task ID"
// @success 200 {object} models.Task
// @Router /task/{task_id} [delete]
func DeleteTaskByID(c *gin.Context) {
	taskID := c.Params.ByName("TaskID")
	var task models.Task
	if err := config.DB.Where("task_id=?", taskID).Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"TaskID" + taskID: "is deleted"})
}
