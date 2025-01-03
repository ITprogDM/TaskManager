package handlers

import (
	"TaskManager/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	task := router.Group("/tasks")
	{
		task.POST("/create", h.CreateTask)
		task.GET("/getall", h.GetAllTasks)
		task.PUT("/update", h.UpdateTask)
		task.DELETE("/:id", h.DeleteTask)
	}

	return router
}
