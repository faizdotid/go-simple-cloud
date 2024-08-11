package main

import (
	"fmt"
	"go-simple-cloud/internal/router"
	"os"

	"github.com/gin-gonic/gin"
	// "go-cloud/app/utils"
)

// Main function
func main() {
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	route := router.NewRoute()
	route.Run(
		fmt.Sprintf(":%s", os.Getenv("PORT")),
	)
}
