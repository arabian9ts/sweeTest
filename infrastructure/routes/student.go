package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
)

func NewStudent(r *gin.RouterGroup, controllers *controllers.RootController, handlers *handler.RootHandler) {
	{
		studentV1 := r.Group("/v1/student")

		studentV1.POST("/login", func(c *gin.Context) { controllers.LoginController.Student.Create(c) })
		studentV1.POST("/sign_up", func(c *gin.Context) { controllers.StudentsController.Create(c) })

		studentV1.Use(handlers.AuthHandler.StudentAuthHandler())

		// me
		{
			me := studentV1.Group("/me")
			me.GET("", handlers.MeHandler.StudentMeHandler())
			me.GET("/lectures", func(c *gin.Context) { controllers.ParticipatedLectureController.Student.Index(c) })
		}

		// lectures
		{
			lectures := studentV1.Group("/lectures")
			lecture := lectures.Group("/:lecture_id")

			lectures.GET("", func(c *gin.Context) { controllers.LecturesController.Index(c) })
			lecture.GET("", func(c *gin.Context) { controllers.LecturesController.Show(c) })

			// tasks
			{
				tasks := lecture.Group("/tasks")
				task := tasks.Group("/:task_id")

				tasks.GET("", func(c *gin.Context) { controllers.TasksController.Index(c) })
				tasks.POST("", func(c *gin.Context) { controllers.TasksController.Create(c) })
				task.PUT("", func(c *gin.Context) { controllers.TasksController.Update(c) })
				task.DELETE("", func(c *gin.Context) { controllers.TasksController.Delete(c) })
			}

			// helps
			{
				helps := lecture.Group("/helps")
				help := helps.Group("/:help_id")

				helps.GET("", func(c *gin.Context) { controllers.HelpsController.Index(c) })
				helps.POST("", func(c *gin.Context) { controllers.HelpsController.Create(c) })
				help.PUT("", func(c *gin.Context) { controllers.HelpsController.Update(c) })
				help.DELETE("", func(c *gin.Context) { controllers.HelpsController.Delete(c) })

				// comments
				{
					comments := help.Group("/comments")
					comment := comments.Group("/:comment_id")

					comments.GET("", func(c *gin.Context) { controllers.CommentsController.Student.Index(c) })
					comments.POST("", func(c *gin.Context) { controllers.CommentsController.Student.Create(c) })
					comment.PUT("", func(c *gin.Context) { controllers.CommentsController.Student.Update(c) })
					comment.DELETE("", func(c *gin.Context) { controllers.CommentsController.Student.Delete(c) })
				}
			}
		}
	}
}
