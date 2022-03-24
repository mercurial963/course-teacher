package server

import (
	"course-teacher/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		teacherGroup := v1.Group("teacher")
		{
			teacher := new(controllers.TeacherController)
			teacherGroup.GET("/:id", teacher.GetByID)
			teacherGroup.POST("/", teacher.Create)
			teacherGroup.PUT("/:id", teacher.UpdateName)
			teacherGroup.DELETE("/:id", teacher.Delete)
		}
		courseGroup := v1.Group("course")
		{
			course := new(controllers.CourseController)
			courseGroup.GET("/:id", course.GetByID)
			courseGroup.POST("/", course.Create)
			courseGroup.PUT("/:id", course.UpdateCourse)
			courseGroup.DELETE("/:id", course.Delete)
		}

	}
	return router

}
