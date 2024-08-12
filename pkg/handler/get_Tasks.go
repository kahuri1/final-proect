package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetTasks(c *gin.Context) {
	search := c.Query("search")
	tasks, err := h.service.GetTasks(search)

	if err != nil {
		log.Printf("Failed GetTasks: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
