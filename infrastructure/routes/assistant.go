package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
)

func NewAssistant(r *gin.RouterGroup, controllers *controllers.RootController, handlers *handler.RootHandler) {
	{
		assistantV1 := r.Group("/v1/assistant")
		assistantV1.POST("/sign_in", func(c *gin.Context) { controllers.LoginController.Assistant.SignIn(c) })
		assistantV1.Use(handlers.AuthHandler.AssistantAuthHandler())

		//
		// me
		//
		{
			me := assistantV1.Group("/me")
			me.GET("", handlers.MeHandler.AssistantMeHandler())
			me.GET("/lectures", func(c *gin.Context) { controllers.ParticipatingLectureController.Assistant.GetLectures(c) })
		}

		//
		// lectures
		//
		{
			lectures := assistantV1.Group("/lectures")
			lecture := lectures.Group("/:lecture_id")

			lectures.GET("", func(c *gin.Context) { controllers.LecturesController.GetLectures(c) })
			lectures.POST("", func(c *gin.Context) { controllers.LecturesController.CreateLecture(c) })
			lecture.GET("", func(c *gin.Context) { controllers.LecturesController.GetLectureById(c) })

			//
			// participation
			//
			{
				participation := lecture.Group("/participation")

				// participate to lecture
				participation.POST("", func(c *gin.Context) { controllers.ParticipationController.AssistantParticipateToLecture(c) })

				// exit from lecture
				participation.DELETE("", func(c *gin.Context) { controllers.ParticipationController.AssistantExitFromLecture(c) })
			}

			//
			// participants
			//
			{
				// get participating students by LectureId
				lecture.GET("/students", func(c *gin.Context) { controllers.ParticipantsController.GetStudentsByLectureId(c) })

				// get participating assistants by LectureId
				lecture.GET("/assistants", func(c *gin.Context) { controllers.ParticipantsController.GetAssistantsByLectureId(c) })

				// get participating teachers by LectureId
				lecture.GET("/teachers", func(c *gin.Context) { controllers.ParticipantsController.GetTeachersByLectureId(c) })
			}

			//
			// helps
			//
			{
				helps := lecture.Group("/helps")
				help := helps.Group("/:help_id")

				helps.GET("", func(c *gin.Context) { controllers.HelpsController.Assistant.GetHelpsByLectureId(c) })
				help.DELETE("", func(c *gin.Context) { controllers.HelpsController.Assistant.DeleteHelp(c) })

				//
				// comments
				//
				{
					comments := help.Group("/comments")
					comment := comments.Group("/:comment_id")

					comments.GET("", func(c *gin.Context) { controllers.CommentsController.Assistant.GetCommentsByHelpId(c) })
					comments.POST("", func(c *gin.Context) { controllers.CommentsController.Assistant.CreateComment(c) })
					comment.PUT("", func(c *gin.Context) { controllers.CommentsController.Assistant.UpdateComment(c) })
					comment.DELETE("", func(c *gin.Context) { controllers.CommentsController.Assistant.DeleteComment(c) })
				}
			}
		}
	}
}
