package infrastructure

import (
	"io"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
	"github.com/arabian9ts/sweeTest/infrastructure/routes"
)

func Router(controllers *controllers.RootController, handlers *handler.RootHandler, version string) (router *gin.Engine) {
	logFile, err := os.OpenFile("./stdout.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile)
	gin.DefaultWriter = io.MultiWriter(logFile)

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	// CORS
	router.Use(cors.Default())

	// to catch panic for slack notification
	router.Use(gin.Recovery())

	api := router.Group("/api")
	api.GET("/alive", handlers.AliveHandler.Echo(version))

	routes.NewStudent(api, controllers, handlers)
	routes.NewAssistant(api, controllers, handlers)
	routes.NewTeacher(api, controllers, handlers)
	routes.NewAdmin(api, controllers, handlers)

	return
}
