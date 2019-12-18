package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
)

func NewTeacher(r *gin.RouterGroup, controllers *controllers.RootController, handlers *handler.RootHandler) {
	{
		teacherV1 := r.Group("/v1/teacher")
		teacherV1.POST("/sign_in", func(c *gin.Context) { controllers.LoginController.Teacher.Create(c) })
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
			teachers.GET("/:teacher_id", func(c *gin.Context) { controllers.TeachersController.GetTeacherById(c) })
			teachers.PUT("/:teacher_id", func(c *gin.Context) { controllers.TeachersController.UpdateTeacher(c) })
		}

		//
		// lectures
		//
		{
			lectures := teacherV1.Group("/lectures")
			lecture := lectures.Group("/:lecture_id")
			lectures.GET("", func(c *gin.Context) { controllers.LecturesController.GetLectures(c) })
			lectures.POST("", func(c *gin.Context) { controllers.LecturesController.CreateLecture(c) })
			lecture.GET("", func(c *gin.Context) { controllers.LecturesController.GetLectureById(c) })
			lecture.PUT("", func(c *gin.Context) { controllers.LecturesController.UpdateLecture(c) })

			//
			// participation
			//
			{
				participation := lecture.Group("/participation")

				// participate to lecture
				participation.POST("", func(c *gin.Context) { controllers.ParticipationController.TeacherParticipateToLecture(c) })

				// exit from lecture
				participation.DELETE("", func(c *gin.Context) { controllers.ParticipationController.TeacherExitFromLecture(c) })
			}

			//
			// participants
			//
			{
				lecture.GET("/students", func(c *gin.Context) { controllers.ParticipantsController.GetStudentsByLectureId(c) })
				lecture.GET("/assistants", func(c *gin.Context) { controllers.ParticipantsController.GetAssistantsByLectureId(c) })
				lecture.GET("/teachers", func(c *gin.Context) { controllers.ParticipantsController.GetTeachersByLectureId(c) })
			}

			//
			// helps
			//
			{
				helps := lecture.Group("/helps")
				help := helps.Group("/:help_id")
				helps.GET("", func(c *gin.Context) { controllers.HelpsController.Teacher.GetHelpsByLectureId(c) })
				help.DELETE("", func(c *gin.Context) { controllers.HelpsController.Teacher.DeleteHelp(c) })

				//
				// comments
				//
				{
					comments := help.Group("/comments")
					comment := comments.Group("/:comment_id")
					comments.GET("", func(c *gin.Context) { controllers.CommentsController.Teacher.Index(c) })
					comments.POST("", func(c *gin.Context) { controllers.CommentsController.Teacher.Create(c) })
					comment.PUT("", func(c *gin.Context) { controllers.CommentsController.Teacher.Update(c) })
					comment.DELETE("", func(c *gin.Context) { controllers.CommentsController.Teacher.Delete(c) })
				}
			}

			//
			// classes
			//
			{
				classes := lecture.Group("/classes")
				class := classes.Group("/:class_id")
				classes.GET("", func(c *gin.Context) { controllers.ClassesController.GetClassesByLectureId(c) })
				classes.POST("", func(c *gin.Context) { controllers.ClassesController.CreateClass(c) })
				class.GET("", func(c *gin.Context) { controllers.ClassesController.GetClassById(c) })
				class.PUT("", func(c *gin.Context) { controllers.ClassesController.UpdateClass(c) })
			}
		}
	}
}
