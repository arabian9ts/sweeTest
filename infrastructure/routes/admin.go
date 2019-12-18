package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
)

func NewAdmin(r *gin.RouterGroup, controllers *controllers.RootController, handlers *handler.RootHandler) {
	{
		adminV1 := r.Group("/v1/admin")
		adminV1.POST("/login", func(c *gin.Context) { controllers.LoginController.Admin.Create(c) })

		adminV1.Use(handlers.AuthHandler.AdminAuthHandler())

		// me
		{
			me := adminV1.Group("/me")
			me.GET("", handlers.MeHandler.AdminMeHandler())
		}

		//  students
		{
			students := adminV1.Group("/students")
			student := students.Group("/:student_id")
			students.GET("", func(c *gin.Context) { controllers.StudentsController.GetStudents(c) })
			students.POST("", func(c *gin.Context) { controllers.StudentsController.SignUp(c) })
			student.GET("", func(c *gin.Context) { controllers.StudentsController.GetStudentById(c) })
			student.PUT("", func(c *gin.Context) { controllers.StudentsController.UpdateStudent(c) })
		}

		// assistants
		{
			assistants := adminV1.Group("/assistants")
			assistant := assistants.Group("/:assistant_id")
			assistants.GET("", func(c *gin.Context) { controllers.AssistantsController.GetAssistants(c) })
			assistants.POST("", func(c *gin.Context) { controllers.AssistantsController.CreateAssistant(c) })
			assistant.GET("", func(c *gin.Context) { controllers.AssistantsController.GetAssistantById(c) })
			assistant.PUT("", func(c *gin.Context) { controllers.AssistantsController.UpdateAssistant(c) })
		}

		// teachers
		{
			teachers := adminV1.Group("/teachers")
			teacher := teachers.Group("/:teacher_id")
			teachers.GET("", func(c *gin.Context) { controllers.TeachersController.GetTeachers(c) })
			teachers.POST("", func(c *gin.Context) { controllers.TeachersController.CreateTeacher(c) })
			teacher.GET("", func(c *gin.Context) { controllers.TeachersController.GetTeacherById(c) })
			teacher.PUT("", func(c *gin.Context) { controllers.TeachersController.UpdateTeacher(c) })
		}

		// lectures
		{
			lectures := adminV1.Group("/lectures")
			lecture := lectures.Group("/:lecture_id")
			lectures.GET("", func(c *gin.Context) { controllers.LecturesController.GetLectures(c) })
			lectures.POST("", func(c *gin.Context) { controllers.LecturesController.CreateLecture(c) })
			lecture.GET("", func(c *gin.Context) { controllers.LecturesController.GetLectureById(c) })
			lecture.PUT("", func(c *gin.Context) { controllers.LecturesController.UpdateLecture(c) })
			lecture.DELETE("", func(c *gin.Context) { controllers.LecturesController.DeleteLecture(c) })
		}
	}
}
