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

	router.GET("/")

	public := router.Group("api")
	{
		public.GET("/endpoints")
		public.POST("/endpoints")
		public.GET("/endpoints:id")
		public.PUT("/endpoints/:id")
		public.DELETE("/endpoints/:id")
	}

	return router
}
