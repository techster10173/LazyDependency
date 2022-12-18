package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"lazydependency/controllers"
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

	apiV1 := router.Group("/api/v1")

	projects := apiV1.Group("/projects")
	{
		projects.GET("/", controllers.GetProjects)
		projects.GET("/:id", controllers.GetProject)
		projects.POST("/", controllers.CreateProject)
		projects.PUT("/:id", controllers.UpdateProject)
		projects.DELETE("/:id", controllers.DeleteProject)
	}

	dependencies := apiV1.Group("/dependencies")
	{
		dependencies.GET("/", controllers.GetDependencies)
		dependencies.GET("/:id", controllers.GetDependency)
		dependencies.POST("/link", controllers.LinkDependency)
	}

	return router
}
