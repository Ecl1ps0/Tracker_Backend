package handler

import (
	"Proctor/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		logrus.Fatalf("Set trusted proxies error: %v\n", err)
		return nil
	}

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

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/profile", h.getProfile)
		}

		tasks := api.Group("/tasks")
		{
			tasks.GET("/:id", h.getTaskByID)
			tasks.POST("/create", h.createTask)
			tasks.DELETE("/delete/:id", h.deleteTask)
		}
	}

	return router
}
