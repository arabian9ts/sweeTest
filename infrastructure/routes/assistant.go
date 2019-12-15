package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
)

func NewAssistant(r *gin.RouterGroup, controllers *controllers.RootController, handlers *handler.RootHandler) {
	{
		assistantV1 := r.Group("/v1/assistant")
		assistantV1.POST("/login", func(c *gin.Context) { controllers.LoginController.Assistant.Create(c) })

		assistantV1.Use(handlers.AuthHandler.AssistantAuthHandler())

		// me
		{
			me := assistantV1.Group("/me")
			me.GET("", handlers.MeHandler.AssistantMeHandler())
			me.GET("/lectures", func(c *gin.Context) { controllers.ParticipatedLectureController.Assistant.Index(c) })
		}

		// lectures
		{
			lectures := assistantV1.Group("/lectures")
			lecture := lectures.Group("/:lecture_id")

			lectures.GET("", func(c *gin.Context) { controllers.LecturesController.Index(c) })
			lectures.POST("", func(c *gin.Context) { controllers.LecturesController.Create(c) })
			lecture.GET("", func(c *gin.Context) { controllers.LecturesController.Show(c) })

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

					comments.GET("", func(c *gin.Context) { controllers.CommentsController.Assistant.Index(c) })
					comments.POST("", func(c *gin.Context) { controllers.CommentsController.Assistant.Create(c) })
					comment.PUT("", func(c *gin.Context) { controllers.CommentsController.Assistant.Update(c) })
					comment.DELETE("", func(c *gin.Context) { controllers.CommentsController.Assistant.Delete(c) })
				}
			}
		}
	}
}
