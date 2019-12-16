package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
)

func NewStudent(r *gin.RouterGroup, controllers *controllers.RootController, handlers *handler.RootHandler) {
	{
		studentV1 := r.Group("/v1/student")

		studentV1.POST("/sign_in", func(c *gin.Context) { controllers.LoginController.Student.LogIn(c) })
		studentV1.POST("/sign_up", func(c *gin.Context) { controllers.StudentsController.SignUp(c) })
		studentV1.Use(handlers.AuthHandler.StudentAuthHandler())

		//
		// me
		//
		{
			me := studentV1.Group("/me")
			me.GET("", handlers.MeHandler.StudentMeHandler())
			me.GET("/lectures", func(c *gin.Context) { controllers.ParticipatingLectureController.Student.GetLectures(c) })
		}

		//
		// lectures
		//
		{
			lectures := studentV1.Group("/lectures")
			lecture := lectures.Group("/:lecture_id")

			lectures.GET("", func(c *gin.Context) { controllers.LecturesController.GetLectures(c) })
			lecture.GET("", func(c *gin.Context) { controllers.LecturesController.GetLectureById(c) })

			//
			// participants
			//
			{
				// get participating students by LectureId
				lecture.GET("/participants", func(c *gin.Context) { controllers.ParticipantsController.GetStudentsByLectureId(c) })
			}

			//
			// participation
			//
			{
				participation := lecture.Group("/participation")

				// participate to lecture
				participation.POST("/participation", func(c *gin.Context) { controllers.ParticipationController.Student.ParticipateToLecture(c) })

				// exit from lecture
				participation.DELETE("/participation", func(c *gin.Context) { controllers.ParticipationController.Student.ExitFromLecture(c) })
			}

			//
			// tasks
			//
			{
				tasks := lecture.Group("/tasks")
				task := tasks.Group("/:task_id")

				tasks.GET("", func(c *gin.Context) { controllers.TasksController.GetTasksByLectureId(c) })
				tasks.POST("", func(c *gin.Context) { controllers.TasksController.CreateTask(c) })
				task.PUT("", func(c *gin.Context) { controllers.TasksController.UpdateTask(c) })
				task.DELETE("", func(c *gin.Context) { controllers.TasksController.DeleteTask(c) })
			}

			//
			// helps
			//
			{
				helps := lecture.Group("/helps")
				help := helps.Group("/:help_id")

				helps.GET("", func(c *gin.Context) { controllers.HelpsController.Student.GetHelpsByLectureId(c) })
				helps.POST("", func(c *gin.Context) { controllers.HelpsController.Student.CreateHelp(c) })
				help.PUT("", func(c *gin.Context) { controllers.HelpsController.Student.UpdateHelp(c) })
				help.DELETE("", func(c *gin.Context) { controllers.HelpsController.Student.DeleteHelp(c) })

				//
				// comments
				//
				{
					comments := help.Group("/comments")
					comment := comments.Group("/:comment_id")

					comments.GET("", func(c *gin.Context) { controllers.CommentsController.Student.GetCommentsByHelpId(c) })
					comments.POST("", func(c *gin.Context) { controllers.CommentsController.Student.CreateComment(c) })
					comment.PUT("", func(c *gin.Context) { controllers.CommentsController.Student.UpdateComment(c) })
					comment.DELETE("", func(c *gin.Context) { controllers.CommentsController.Student.DeleteComment(c) })
				}
			}
		}
	}
}
