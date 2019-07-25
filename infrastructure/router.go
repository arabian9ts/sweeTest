package infrastructure

import (
	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(root *controllers.RootController) (router *gin.Engine) {
	router = gin.Default()
	router.Use(cors.Default())

	router.GET("/lectures", func(c *gin.Context) { root.LecturesController.Index(c) })
	router.POST("/lectures", func(c *gin.Context) { root.LecturesController.Create(c) })

	lecture := router.Group("/lectures/:lecture_id")
	{
		lecture.GET("/", func(c *gin.Context) { root.LecturesController.Show(c) })
		lecture.PUT("/", func(c *gin.Context) { root.LecturesController.Update(c) })
		lecture.DELETE("/", func(c *gin.Context) { root.LecturesController.Delete(c) })
		lecture.GET("/tasks", func(c *gin.Context) { root.TasksController.Index(c) })
		lecture.POST("/tasks", func(c *gin.Context) { root.TasksController.Create(c) })

		task := lecture.Group("/tasks/:task_id")
		{
			task.PUT("/", func(c *gin.Context) { root.TasksController.Update(c) })
			task.DELETE("/", func(c *gin.Context) { root.TasksController.Delete(c) })
		}
	}

	router.POST("/students", func(c *gin.Context) { root.StudentsController.Create(c) })
	router.GET("/students/:id", func(c *gin.Context) { root.StudentsController.Show(c) })

	router.POST("/tas", func(c *gin.Context) { root.TasController.Create(c) })
	router.GET("/tas/:id", func(c *gin.Context) { root.TasController.Show(c) })

	router.POST("/teachers", func(c *gin.Context) { root.TeachersController.Create(c) })
	router.GET("/teachers/:id", func(c *gin.Context) { root.TeachersController.Show(c) })

	return
}
