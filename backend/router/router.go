package router

import (
	"lazydependency/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer(prod *bool) (r *gin.Engine) {
	if *prod {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	// TODO Setup CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	apiV1 := router.Group("/api/v1")

	dependencies := apiV1.Group("/dependencies")
	{
		dependencies.POST("/upload", controllers.UploadConnections)
	}

	return router
}
