package handlers

import (
	"TaskManager/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.CreatedAt = time.Now()
	task.UpdateAt = time.Now()

	_, err := h.service.CreateTask(c.Request.Context(), task.Title, task.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.UpdateAt = time.Now()
	err := h.service.UpdateTask(c.Request.Context(), task.ID, task.Title, task.Description, task.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request task ID"})
		return
	}
	err = h.service.DeleteTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
