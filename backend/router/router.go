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

	users := apiV1.Group("/users")
	{
		users.POST("/login", controllers.Login)
	}

	project := apiV1.Group("/projects")
	{
		project.POST("/", controllers.CreateProject)
		project.PATCH("/:id", controllers.UpdateProject)
		project.DELETE("/:id", controllers.DeleteProject)
		project.GET("/", controllers.ListProjects)
		project.GET("/:id", controllers.GetProject)
		project.GET("/reccomendations", controllers.GetReccomendations)
	}

	dependencies := apiV1.Group("/dependencies")
	{
		dependencies.GET("/:id", controllers.GetDependency)
		dependencies.POST("/like", controllers.LikeDependency)
		dependencies.POST("/unlike", controllers.UnlikeDependency)
		dependencies.POST("/comment", controllers.CommentDependency)
	}

	return router
}
