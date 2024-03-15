package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-speech/internal/config"
)

func StartServer(handlers *config.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12,
	}))

	public := router.Group("api")
	{
		public.POST("/user", handlers.UserHandler.CreateUser)
		public.GET("/user/:uuid", handlers.UserHandler.GetUser)
		public.POST("/verify-user", handlers.UserHandler.VerifyUser)
		public.GET("/roles", handlers.UserHandler.GetRoles)
		public.GET("/department", handlers.MasterHandler.GetAllDepartment)
	}

	admin := router.Group("api")
	{
		admin.POST("/class", handlers.MasterHandler.CreateClass)
		admin.POST("/department", handlers.MasterHandler.CreateDepartment)
		admin.GET("/department/:uuid", handlers.MasterHandler.GetDepartment)
		admin.PUT("/department/:uuid", handlers.MasterHandler.UpdateDepartment)
		admin.DELETE("/department/:uuid", handlers.MasterHandler.DeleteDepartment)
	}

	return router
}
