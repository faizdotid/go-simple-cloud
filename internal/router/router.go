package router

import (
	"go-simple-cloud/internal/controllers"
	"go-simple-cloud/internal/controllers/files"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB) *gin.Engine {

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
	r.Static("/assets/static", "assets/static")

	apiGroup := r.Group("/api/v1") // new api group

	fileController := files.NewFileController(db)
	apiGroup.GET("/", controllers.IndexController)
	apiGroup.POST("/files", fileController.Upload)
	apiGroup.GET("/files", fileController.Index)
	apiGroup.GET("/file/:url", fileController.Show)
	apiGroup.GET("/file/:url/download", fileController.Download)

	return r
}
