package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-speech/internal/config"
	"github.com/iki-rumondor/go-speech/internal/middleware"
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
		public.GET("/public/classes", handlers.UserHandler.GetAllClasses)
		public.GET("/roles", handlers.UserHandler.GetRoles)
		public.GET("/department", handlers.MasterHandler.GetAllDepartment)
	}

	admin := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("ADMIN"))
	{
		admin.POST("/department", handlers.MasterHandler.CreateDepartment)
		admin.GET("/department/:uuid", handlers.MasterHandler.GetDepartment)
		admin.PUT("/department/:uuid", handlers.MasterHandler.UpdateDepartment)
		admin.DELETE("/department/:uuid", handlers.MasterHandler.DeleteDepartment)

		admin.GET("/teachers", handlers.UserHandler.GetTeachers)
		admin.PATCH("/teacher/:uuid/activate", handlers.UserHandler.ActivateUser)
	}

	teacher := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("DOSEN"), middleware.SetUserUuid())
	{
		teacher.POST("/classes", handlers.MasterHandler.CreateClass)
		teacher.GET("/classes", handlers.MasterHandler.GetClasses)
		teacher.GET("/classes/:uuid", handlers.MasterHandler.GetClass)
		teacher.PUT("/classes/:uuid", handlers.MasterHandler.UpdateClass)
		teacher.DELETE("/classes/:uuid", handlers.MasterHandler.DeleteClass)
		teacher.GET("/classes/request", handlers.UserHandler.GetRequestClasses)
		teacher.PATCH("/classes/:uuid/request", handlers.UserHandler.UpdateStatusClassReq)

		teacher.POST("/videos", handlers.FileHandler.CreateVideo)
		teacher.GET("/videos/classes/:uuid", handlers.FileHandler.GetClassVideos)
	}

	student := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("MAHASISWA"), middleware.SetUserUuid())
	{
		student.POST("/class/register", handlers.UserHandler.CreateClassRequest)
		student.GET("/class/request/students", handlers.UserHandler.GetStudentRequestClasses)
	}

	return router
}
