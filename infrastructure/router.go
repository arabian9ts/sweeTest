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
		studentEndPoint.POST("/login", func(c *gin.Context) { controllers.LoginController.Student.Create(c) })

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

		studentEndPoint.GET("/lectures/:lecture_id/helps/:help_id/comments", func(c *gin.Context) { controllers.CommentsController.Student.Index(c) })
		studentEndPoint.POST("/lectures/:lecture_id/helps/:help_id/comments", func(c *gin.Context) { controllers.CommentsController.Student.Create(c) })
		studentEndPoint.PUT("/lectures/:lecture_id/helps/:help_id/comments/:comment_id", func(c *gin.Context) { controllers.CommentsController.Student.Update(c) })
		studentEndPoint.DELETE("/lectures/:lecture_id/helps/:help_id/comments/:comment_id", func(c *gin.Context) { controllers.CommentsController.Student.Delete(c) })

		studentEndPoint.GET("/lectures/:lecture_id/students/:student_id/participation", func(c *gin.Context) { controllers.ParticipationController.Student.Create(c) })
		studentEndPoint.DELETE("/lectures/:lecture_id/students/:student_id/participation/", func(c *gin.Context) { controllers.ParticipationController.Student.Delete(c) })

		studentEndPoint.GET("/me/lectures/", func(c *gin.Context) { controllers.ParticipatedLectureController.Student.Index(c) })
	}

	assistantEndPoint := v1APIEndpoint.Group("/assistant")
	{
		assistantEndPoint.POST("/login", func(c *gin.Context) { controllers.LoginController.Assistant.Create(c) })

		assistantEndPoint.Use(handlers.AuthHandler.AssistantAuthHandler())
		assistantEndPoint.GET("/me", handlers.MeHandler.AssistantMeHandler())
		assistantEndPoint.POST("/", func(c *gin.Context) { controllers.AssistantsController.Create(c) })
		assistantEndPoint.GET("/assistants/:assistant_id", func(c *gin.Context) { controllers.AssistantsController.Show(c) })
		assistantEndPoint.PUT("/assistants/:assistant_id", func(c *gin.Context) { controllers.AssistantsController.Update(c) })

		assistantEndPoint.GET("/lectures/:lecture_id/helps", func(c *gin.Context) { controllers.HelpsController.Index(c) })
		assistantEndPoint.POST("/lectures/:lecture_id/helps", func(c *gin.Context) { controllers.HelpsController.Create(c) })
		assistantEndPoint.PUT("/lectures/:lecture_id/helps/:help_id", func(c *gin.Context) { controllers.HelpsController.Update(c) })
		assistantEndPoint.DELETE("/lectures/:lecture_id/helps/:help_id", func(c *gin.Context) { controllers.HelpsController.Delete(c) })

		assistantEndPoint.GET("/lectures/:lecture_id/helps/:help_id/comments", func(c *gin.Context) { controllers.CommentsController.Assistant.Index(c) })
		assistantEndPoint.POST("/lectures/:lecture_id/helps/:help_id/comments", func(c *gin.Context) { controllers.CommentsController.Assistant.Create(c) })
		assistantEndPoint.PUT("/lectures/:lecture_id/helps/:help_id/comments/:comment_id", func(c *gin.Context) { controllers.CommentsController.Assistant.Update(c) })
		assistantEndPoint.DELETE("/lectures/:lecture_id/helps/:help_id/comments/:comment_id", func(c *gin.Context) { controllers.CommentsController.Assistant.Delete(c) })

		assistantEndPoint.GET("/lectures/:lecture_id/assistants/:student_id/participation", func(c *gin.Context) { controllers.ParticipationController.Assistant.Create(c) })
		assistantEndPoint.DELETE("/lectures/:lecture_id/assistants/:student_id/participation/", func(c *gin.Context) { controllers.ParticipationController.Assistant.Delete(c) })

		assistantEndPoint.GET("/me/lectures", func(c *gin.Context) { controllers.ParticipatedLectureController.Assistant.Index(c) })

		assistantEndPoint.GET("lectures/:lecture_id/participants/students", func(c *gin.Context) { controllers.ParticipantsController.Student.Index(c) })
		assistantEndPoint.GET("lectures/:lecture_id/participants/assistants", func(c *gin.Context) { controllers.ParticipantsController.Assistant.Index(c) })
		assistantEndPoint.GET("lectures/:lecture_id/participants/teachers", func(c *gin.Context) { controllers.ParticipantsController.Teacher.Index(c) })
	}

	teacherEndPoint := v1APIEndpoint.Group("/teacher")
	{
		teacherEndPoint.POST("/login", func(c *gin.Context) { controllers.LoginController.Teacher.Create(c) })

		teacherEndPoint.Use(handlers.AuthHandler.TeacherAuthHandler())
		teacherEndPoint.GET("/me", handlers.MeHandler.TeacherMeHandler())
		teacherEndPoint.POST("/", func(c *gin.Context) { controllers.TeachersController.Create(c) })
		teacherEndPoint.GET("/teachers/:teacher_id", func(c *gin.Context) { controllers.TeachersController.Show(c) })
		teacherEndPoint.PUT("/teachers/:teacher_id", func(c *gin.Context) { controllers.TeachersController.Update(c) })

		teacherEndPoint.GET("/lectures/:lecture_id/helps/:help_id/comments", func(c *gin.Context) { controllers.CommentsController.Teacher.Index(c) })
		teacherEndPoint.POST("/lectures/:lecture_id/helps/:help_id/comments", func(c *gin.Context) { controllers.CommentsController.Teacher.Create(c) })
		teacherEndPoint.PUT("/lectures/:lecture_id/helps/:help_id/comments/:comment_id", func(c *gin.Context) { controllers.CommentsController.Teacher.Update(c) })
		teacherEndPoint.DELETE("/lectures/:lecture_id/helps/:help_id/comments/:comment_id", func(c *gin.Context) { controllers.CommentsController.Teacher.Delete(c) })

		teacherEndPoint.GET("/lectures/:lecture_id/teachers/:teacher_id/participation", func(c *gin.Context) { controllers.ParticipationController.Teacher.Create(c) })
		teacherEndPoint.DELETE("/lectures/:lecture_id/teachers/:teacher_id/participation/", func(c *gin.Context) { controllers.ParticipationController.Teacher.Delete(c) })

		teacherEndPoint.GET("/me/lectures/", func(c *gin.Context) { controllers.ParticipatedLectureController.Teacher.Index(c) })

		teacherEndPoint.GET("/lectures/:lecture_id/participants/students", func(c *gin.Context) { controllers.ParticipantsController.Student.Index(c) })
		teacherEndPoint.GET("/lectures/:lecture_id/participants/assistants", func(c *gin.Context) { controllers.ParticipantsController.Assistant.Index(c) })
		teacherEndPoint.GET("/lectures/:lecture_id/participants/teachers", func(c *gin.Context) { controllers.ParticipantsController.Teacher.Index(c) })
	}

	adminEndPoint := v1APIEndpoint.Group("/admin")
	{
		adminEndPoint.POST("/login", func(c *gin.Context) { controllers.LoginController.Admin.Create(c) })

		adminEndPoint.Use(handlers.AuthHandler.AdminAuthHandler())
		adminEndPoint.GET("/me", handlers.MeHandler.AdminMeHandler())

		adminEndPoint.GET("/lectures", func(c *gin.Context) { controllers.LecturesController.Index(c) })
		adminEndPoint.POST("/lectures", func(c *gin.Context) { controllers.LecturesController.Create(c) })

		adminEndPoint.GET("/teachers", func(c *gin.Context) { controllers.TeachersController.Index(c) })
		adminEndPoint.POST("/teachers", func(c *gin.Context) { controllers.TeachersController.Create(c) })
		adminEndPoint.GET("/teachers/:teacher_id", func(c *gin.Context) { controllers.TeachersController.Show(c) })
		adminEndPoint.PUT("/teachers/:teacher_id", func(c *gin.Context) { controllers.TeachersController.Update(c) })
	}

	return
}
