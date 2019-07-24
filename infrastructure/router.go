package infrastructure

import (
	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/gin-gonic/gin"
)


func Router(root *controllers.RootController) (router *gin.Engine) {
	router = gin.Default()

	router.GET("/lectures", func(c *gin.Context) {root.LecturesController.Index(c)})
	router.POST("/lectures", func(c *gin.Context) {root.LecturesController.Create(c)})
	router.GET("/lectures/:id", func(c *gin.Context) {root.LecturesController.Show(c)})
	router.PUT("/lectures/:id", func(c *gin.Context) {root.LecturesController.Update(c)})
	router.DELETE("/lectures/:id", func(c *gin.Context) {root.LecturesController.Delete(c)})

	router.POST("/students", func(c *gin.Context) {root.StudentsController.Create(c)})
	router.GET("/students/:id", func(c *gin.Context) {root.StudentsController.Show(c)})

	router.POST("/tas", func(c *gin.Context) {root.TasController.Create(c)})
	router.GET("/tas/:id", func(c *gin.Context) {root.TasController.Show(c)})

	router.POST("/teachers", func(c *gin.Context) {root.TeachersController.Create(c)})
	router.GET("/teachers/:id", func(c *gin.Context) {root.TeachersController.Show(c)})

	return
}