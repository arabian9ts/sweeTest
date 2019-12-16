package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
)

func NewTeacher(r *gin.RouterGroup, controllers *controllers.RootController, handlers *handler.RootHandler) {
	{
		teacherV1 := r.Group("/v1/teacher")
		teacherV1.POST("/login", func(c *gin.Context) { controllers.LoginController.Teacher.Create(c) })
		teacherV1.Use(handlers.AuthHandler.TeacherAuthHandler())

		//
		// me
		//
		{
			me := teacherV1.Group("/me")
			me.GET("", handlers.MeHandler.TeacherMeHandler())
			me.GET("/lectures", func(c *gin.Context) { controllers.ParticipatingLectureController.Teacher.GetLectures(c) })
		}

		//
		// teachers
		//
		{
			teachers := teacherV1.Group("/teachers")
			teachers.GET("/:teacher_id", func(c *gin.Context) { controllers.TeachersController.Show(c) })
			teachers.PUT("/:teacher_id", func(c *gin.Context) { controllers.TeachersController.Update(c) })
		}

		//
		// lectures
		//
		{
			lectures := teacherV1.Group("/lectures")
			lecture := lectures.Group("/:lecture_id")
			lectures.GET("", func(c *gin.Context) { controllers.LecturesController.Index(c) })
			lectures.POST("", func(c *gin.Context) { controllers.LecturesController.Create(c) })
			lecture.GET("", func(c *gin.Context) { controllers.LecturesController.Show(c) })
			lecture.PUT("", func(c *gin.Context) { controllers.LecturesController.Update(c) })

			//
			// participation
			//
			{
				participation := lecture.Group("/participation")

				// participate to lecture
				participation.POST("", func(c *gin.Context) { controllers.ParticipationController.Teacher.ParticipateToLecture(c) })

				// exit from lecture
				participation.DELETE("", func(c *gin.Context) { controllers.ParticipationController.Teacher.ExitFromLecture(c) })
			}

			//
			// participants
			//
			{
				lecture.GET("/students", func(c *gin.Context) { controllers.ParticipantsController.Teacher.GetStudentsByLectureId(c) })
				lecture.GET("/assistants", func(c *gin.Context) { controllers.ParticipantsController.Teacher.GetAssistantsByLectureId(c) })
				lecture.GET("/teachers", func(c *gin.Context) { controllers.ParticipantsController.Teacher.GetTeachersByLectureId(c) })
			}

			//
			// helps
			//
			{
				helps := lecture.Group("/helps")
				help := helps.Group("/:help_id")
				helps.GET("", func(c *gin.Context) { controllers.HelpsController.Index(c) })
				helps.POST("", func(c *gin.Context) { controllers.HelpsController.Create(c) })
				help.PUT("", func(c *gin.Context) { controllers.HelpsController.Update(c) })
				help.DELETE("", func(c *gin.Context) { controllers.HelpsController.Delete(c) })

				//
				// comments
				//
				{
					comments := help.Group("/comments")
					comment := comments.Group("/:comment_id")
					comments.GET("", func(c *gin.Context) { controllers.CommentsController.Teacher.Index(c) })
					comments.POST("", func(c *gin.Context) { controllers.CommentsController.Teacher.Create(c) })
					comment.PUT("/:comment_id", func(c *gin.Context) { controllers.CommentsController.Teacher.Update(c) })
					comment.DELETE("/:comment_id", func(c *gin.Context) { controllers.CommentsController.Teacher.Delete(c) })
				}
			}
		}
	}
}
