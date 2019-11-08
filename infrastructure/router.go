package infrastructure

import (
	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(controllers *controllers.RootController, handlers *handler.RootHandler) (router *gin.Engine) {
	router = gin.Default()
	router.Use(cors.Default())

	v1APIEndpoint := router.Group("/api/v1")

	studentEndPoint := v1APIEndpoint.Group("/student")
	{
		studentEndPoint.POST("/login", func(c *gin.Context) { controllers.StudentLoginController.Create(c) })

		studentEndPoint.Use(handlers.AuthHandler.StudentAuthHandler())
		studentEndPoint.GET("/me", handlers.MeHandler.StudentMeHandler())
		studentEndPoint.POST("/", func(c *gin.Context) { controllers.StudentsController.Create(c) })
		studentEndPoint.GET("/students/:student_id", func(c *gin.Context) { controllers.StudentsController.Show(c) })
		studentEndPoint.PUT("/students/:student_id", func(c *gin.Context) { controllers.StudentsController.Update(c) })
		studentEndPoint.GET("/lectures", func(c *gin.Context) { controllers.LecturesController.Index(c) })
		studentEndPoint.POST("/lectures", func(c *gin.Context) { controllers.LecturesController.Create(c) })

		studentEndPoint.GET("/lectures/:lecture_id", func(c *gin.Context) { controllers.LecturesController.Show(c) })
		studentEndPoint.PUT("/lectures/:lecture_id", func(c *gin.Context) { controllers.LecturesController.Update(c) })
		studentEndPoint.DELETE("/lectures/:lecture_id", func(c *gin.Context) { controllers.LecturesController.Delete(c) })

		studentEndPoint.GET("/lectures/:lecture_id/tasks", func(c *gin.Context) { controllers.TasksController.Index(c) })
		studentEndPoint.POST("/lectures/:lecture_id/tasks", func(c *gin.Context) { controllers.TasksController.Create(c) })
		studentEndPoint.PUT("/lectures/:lecture_id/tasks/:task_id", func(c *gin.Context) { controllers.TasksController.Update(c) })
		studentEndPoint.DELETE("/lectures/:lecture_id/tasks/:task_id", func(c *gin.Context) { controllers.TasksController.Delete(c) })

		studentEndPoint.GET("/lectures/:lecture_id/helps", func(c *gin.Context) { controllers.HelpsController.Index(c) })
		studentEndPoint.POST("/lectures/:lecture_id/helps", func(c *gin.Context) { controllers.HelpsController.Create(c) })
		studentEndPoint.PUT("/lectures/:lecture_id/helps/:help_id", func(c *gin.Context) { controllers.HelpsController.Update(c) })
		studentEndPoint.DELETE("/lectures/:lecture_id/helps/:help_id", func(c *gin.Context) { controllers.HelpsController.Delete(c) })

		studentEndPoint.GET("/lectures/:lecture_id/helps/:help_id/comments", func(c *gin.Context) { controllers.StudentCommentsController.Index(c) })
		studentEndPoint.POST("/lectures/:lecture_id/helps/:help_id/comments", func(c *gin.Context) { controllers.StudentCommentsController.Create(c) })
		studentEndPoint.PUT("/lectures/:lecture_id/helps/:help_id/comments/:comment_id", func(c *gin.Context) { controllers.StudentCommentsController.Update(c) })
		studentEndPoint.DELETE("/lectures/:lecture_id/helps/:help_id/comments/:comment_id", func(c *gin.Context) { controllers.StudentCommentsController.Delete(c) })
	}

	assistantEndPoint := v1APIEndpoint.Group("/assistant")
	{
		assistantEndPoint.POST("/login", func(c *gin.Context) { controllers.AssistantLoginController.Create(c) })

		assistantEndPoint.Use(handlers.AuthHandler.AssistantAuthHandler())
		assistantEndPoint.GET("/me", handlers.MeHandler.AssistantMeHandler())
		assistantEndPoint.POST("/", func(c *gin.Context) { controllers.AssistantsController.Create(c) })
		assistantEndPoint.GET("/assistants/:assistant_id", func(c *gin.Context) { controllers.AssistantsController.Show(c) })
		assistantEndPoint.PUT("/assistants/:assistant_id", func(c *gin.Context) { controllers.AssistantsController.Update(c) })
	}

	teacherEndPoint := v1APIEndpoint.Group("/teacher")
	{
		teacherEndPoint.POST("/login", func(c *gin.Context) { controllers.TeacherLoginController.Create(c) })

		teacherEndPoint.Use(handlers.AuthHandler.TeacherAuthHandler())
		teacherEndPoint.GET("/me", handlers.MeHandler.TeacherMeHandler())
		teacherEndPoint.POST("/", func(c *gin.Context) { controllers.TeachersController.Create(c) })
		teacherEndPoint.GET("/teachers/:teacher_id", func(c *gin.Context) { controllers.TeachersController.Show(c) })
		teacherEndPoint.PUT("/teachers/:teacher_id", func(c *gin.Context) { controllers.TeachersController.Update(c) })
	}

	adminEndPoint := v1APIEndpoint.Group("/admin")
	{
		adminEndPoint.POST("/login", func(c *gin.Context) { controllers.AdminLoginController.Create(c) })

		adminEndPoint.Use(handlers.AuthHandler.AdminAuthHandler())
		adminEndPoint.GET("/me", handlers.MeHandler.AdminMeHandler())
	}

	return
}
