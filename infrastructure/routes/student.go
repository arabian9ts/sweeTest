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
				participation.POST("", func(c *gin.Context) { controllers.ParticipationController.StudentParticipateToLecture(c) })

				// exit from lecture
				participation.DELETE("", func(c *gin.Context) { controllers.ParticipationController.StudentExitFromLecture(c) })
			}

			//
			// tasks
			//
			{
				tasks := lecture.Group("/tasks")
				task := tasks.Group("/:task_id")

				tasks.GET("", func(c *gin.Context) { controllers.TasksController.GetTasksByClassId(c) })
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

			//
			// classes
			//
			{
				classes := lecture.Group("/classes")
				class := classes.Group("/:class_id")
				classes.GET("", func(c *gin.Context) { controllers.ClassesController.GetClassesByLectureId(c) })
				class.GET("", func(c *gin.Context) { controllers.ClassesController.GetClassById(c) })

				//
				// tasks
				//
				{
					tasks := class.Group("/tasks")
					task := tasks.Group("/:task_id")
					tasks.GET("", func(c *gin.Context) { controllers.TasksController.GetTasksByClassId(c) })

					//
					// submissions
					//
					{
						submissions := task.Group("/submissions")
						submission := submissions.Group("/:submission_id")
						submissions.POST("", func(c *gin.Context) { controllers.SubmissionsController.CreateSubmission(c) })
						//
						// submissionTexts
						//
						{
							submissionTexts := submission.Group("/submissionTexts")
							submissionText := submissionTexts.Group("/:submissionText_id")
							submissionTexts.GET("", func(c *gin.Context) { controllers.SubmissionsController.GetSubmissionTextsBySubmissionID(c) })
							submissionTexts.POST("", func(c *gin.Context) { controllers.SubmissionsController.CreateSubmissionText(c) })
							submissionText.PUT("", func(c *gin.Context) { controllers.SubmissionsController.UpdateSubmissionText(c) })
							submissionText.DELETE("", func(c *gin.Context) { controllers.SubmissionsController.DeleteSubmissionText(c) })
						}
					}
				}
			}
		}
	}
}
