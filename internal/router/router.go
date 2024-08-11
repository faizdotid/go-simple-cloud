package router

import (
	"go-simple-cloud/internal/controllers"
	"go-simple-cloud/internal/services"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {

	dbService, err := services.NewDatabaseService()
	if err != nil {
		panic(err)
	}

	r := gin.Default(func(e *gin.Engine) {
		e.Use(gin.Logger())
		e.Use(func(ctx *gin.Context) {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
			ctx.Next()
		})
	})
	r.StaticFile("/", "web/index.html")

	apiGroup := r.Group("/api/v1") // new api group

	fileController := controllers.NewFileController(dbService)
	apiGroup.GET("/", controllers.IndexController)
	apiGroup.POST("/files", fileController.Upload)
	apiGroup.GET("/files", fileController.Index)
	apiGroup.GET("/file/:url", fileController.Show)
	apiGroup.GET("/file/:url/download", fileController.Download)

	return r
}
