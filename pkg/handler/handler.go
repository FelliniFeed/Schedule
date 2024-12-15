package handler

import (
	"schedule/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	class := router.Group("/class")
	{
		class.GET("/", h.getClasses)
		class.POST("/", h.createClass)
	}

	student := router.Group("/student")
	{
		student.GET("/", h.getStudents)
		student.POST("/", h.createStudent)
	}

	return router
}