package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/final-proect/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) TaskDone(c *gin.Context) {

	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}
	task, err := h.service.GetTaskById(id)

	if task.Repeat == "" {
		err = h.service.DeleteTask(id)
		if err != nil {
			log.Printf("Failed to NextDate: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	} else {
		task.Date, err = usecase.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			log.Printf("Failed to NextDate: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		_, err = h.service.TaskDone(task)
		if err != nil {
			log.Printf("Failed to create task: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		log.Info("изменение дня")
	}

}
