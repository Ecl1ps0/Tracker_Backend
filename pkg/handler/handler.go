package handler

import (
	"Proctor/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signin", h.singIn)
		auth.POST("/signup", h.singUp)
	}

	messageTransfer := router.Group("/connection")
	{
		messageTransfer.GET("/ws", h.webSocket)
		messageTransfer.POST("/fileUpload", h.fileUpload)
	}

	return router
}
