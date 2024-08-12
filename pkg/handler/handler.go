package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/final-proect/pkg/model"
	"github.com/spf13/viper"
)

type todoService interface {
	CreateTask(t model.Task) (int64, error)
}

type Handler struct {
	service todoService
}

func Newhandler(service todoService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	webDir := viper.GetString("dir")

	router := gin.New()

	router.GET("/api/nextdate", h.NextDateHandler)
	taskGroup := router.Group("/api/task")
	{
		taskGroup.POST("/", h.CreateTask)
	}
	// Обработка статических файлов
	router.Static("/js", webDir+"/js")                       // Обслуживание JS файлов
	router.Static("/css", webDir+"/css")                     // Обслуживание CSS файлов
	router.StaticFile("/favicon.ico", webDir+"/favicon.ico") // Обслуживание favicon.ico
	router.StaticFile("/", webDir+"/index.html")             // Обслуживание index.html по корневому пути

	return router
}
