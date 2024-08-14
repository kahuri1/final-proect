package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/final-proect/pkg/model"
	"github.com/spf13/viper"
)

type todoService interface {
	CreateTask(t model.Task) (int64, error)
	GetTasks(search string) (model.TasksResp, error)
	GetTaskById(id int) (*model.Task, error)
	UpdateTask(task *model.Task) (bool, error)
	TaskDone(task *model.Task) (bool, error)
	DeleteTask(id int) error
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
	router.GET("/api/tasks", h.GetTasks)
	taskGroup := router.Group("/api/task")
	{
		taskGroup.POST("", h.CreateTask)
		taskGroup.GET("", h.GetTaskID)
		taskGroup.PUT("", h.UpdateTask)
		taskGroup.POST("/done", h.TaskDone)
		taskGroup.DELETE("", h.DeleteTask)
	}
	//Обработка статических файлов
	router.Static("/js", webDir+"/js")                       // Обслуживание JS файлов
	router.Static("/css", webDir+"/css")                     // Обслуживание CSS файлов
	router.StaticFile("/favicon.ico", webDir+"/favicon.ico") // Обслуживание index.html по корневому пути
	router.NoRoute(func(c *gin.Context) {
		c.File(webDir + "/index.html")
	})
	return router
}
