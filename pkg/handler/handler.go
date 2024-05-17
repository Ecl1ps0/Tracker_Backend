package handler

import (
	"Proctor/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.Static("/docs", "./docs")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/docs/swagger.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	auth := router.Group("/auth")
	{
		auth.POST("/signin", h.singIn)
		auth.POST("/signup", h.singUp)
	}

	connection := router.Group("/connection")
	{
		connection.GET("/ws", h.webSocket)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/profile", h.getProfile)
			users.POST("/:studentID/add-to-task/:taskID", h.addStudentToTask)
		}

		tasks := api.Group("/tasks")
		{
			tasks.GET("/", h.getAllTasks)
			tasks.GET("/:id", h.getTaskByID)
			tasks.GET("/teacher/:id", h.getAllTeacherTasks)
			tasks.GET("/student/:id", h.getAllStudentTasks)
			tasks.POST("/create", h.createTask)
			tasks.PUT("/update/:id", h.updateTask)
			tasks.DELETE("/delete/:id", h.deleteTask)
		}

		solutions := api.Group("/solutions")
		{
			solutions.GET("/", h.getAllSolutions)
			solutions.POST("/on-student-task/:id", h.createSolution)
			solutions.PUT("/update-cheating-rate/:id", h.updateSolutionCheatingRate)
			solutions.PUT("/update-final-grade/:id", h.updateSolutionFinalGrade)
		}

		reports := api.Group("/reports")
		{
			reports.GET("/", h.getAllReports)
			reports.POST("/createReport", h.createReport)
		}
	}

	return router
}
