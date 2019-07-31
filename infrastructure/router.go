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

	studentEndPoint := router.Group("/student")
	{
		studentEndPoint.POST("/login", func(c *gin.Context) { controllers.StudentLoginController.Create(c) })
		studentEndPoint.Use(handlers.AuthHandler.StudentAuthHandler())
		{
			studentEndPoint.POST("/", func(c *gin.Context) { controllers.StudentsController.Create(c) })
			studentEndPoint.GET("/me", handlers.MeHandler.StudentMeHandler())
			student := studentEndPoint.Group("/students/:student_id")
			{
				student.GET("/", func(c *gin.Context) { controllers.StudentsController.Show(c) })
			}
		}
	}

	assistantEndPoint := router.Group("/assistant")
	{
		assistantEndPoint.POST("/login", func(c *gin.Context) { controllers.AssistantLoginController.Create(c) })
		assistantEndPoint.Use(handlers.AuthHandler.AssistantAuthHandler())
		{
			assistantEndPoint.POST("/", func(c *gin.Context) { controllers.AssistantsController.Create(c) })
			assistantEndPoint.GET("/me", handlers.MeHandler.AssistantMeHandler())
			assistant := assistantEndPoint.Group("/assistants/:assistant_id")
			{
				assistant.GET("/", func(c *gin.Context) { controllers.AssistantsController.Show(c) })
			}
		}
	}

	teacherEndPoint := router.Group("/teacher")
	{
		teacherEndPoint.POST("/login", func(c *gin.Context) { controllers.TeacherLoginController.Create(c) })
		teacherEndPoint.Use(handlers.AuthHandler.TeacherAuthHandler())
		{
			teacherEndPoint.POST("/", func(c *gin.Context) { controllers.TeachersController.Create(c) })
			teacherEndPoint.GET("/me", handlers.MeHandler.TeacherMeHandler())
			teacher := teacherEndPoint.Group("/teachers/:teacher_id")
			{
				teacher.GET("/", func(c *gin.Context) { controllers.TeachersController.Show(c) })
			}
		}
	}

	adminEndPoint := router.Group("/admins")
	{
		adminEndPoint.POST("/login", func(c *gin.Context) { controllers.AdminLoginController.Create(c) })
		adminEndPoint.Use(handlers.AuthHandler.AdminAuthHandler())
		{
			adminEndPoint.GET("/me", handlers.MeHandler.AdminMeHandler())
		}
	}

	router.GET("/lectures", func(c *gin.Context) { controllers.LecturesController.Index(c) })
	router.POST("/lectures", func(c *gin.Context) { controllers.LecturesController.Create(c) })

	lecture := router.Group("/lectures/:lecture_id")
	{
		lecture.GET("/", func(c *gin.Context) { controllers.LecturesController.Show(c) })
		lecture.PUT("/", func(c *gin.Context) { controllers.LecturesController.Update(c) })
		lecture.DELETE("/", func(c *gin.Context) { controllers.LecturesController.Delete(c) })
		lecture.GET("/tasks", func(c *gin.Context) { controllers.TasksController.Index(c) })
		lecture.POST("/tasks", func(c *gin.Context) { controllers.TasksController.Create(c) })

		task := lecture.Group("/tasks/:task_id")
		{
			task.PUT("/", func(c *gin.Context) { controllers.TasksController.Update(c) })
			task.DELETE("/", func(c *gin.Context) { controllers.TasksController.Delete(c) })
		}
	}

	return
}
