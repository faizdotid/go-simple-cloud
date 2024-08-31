package main

import (
	"go-simple-cloud/internal/router"
	"go-simple-cloud/internal/services"
	"go-simple-cloud/pkg/schedule"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func bootstrap(port string, mode string) {
	// set gin mode
	gin.SetMode(mode)

	// init database service
	dbService, err := services.NewDatabaseService()
	if err != nil {
		log.Fatalf("Failed to initialize database service: %v", err)
	}

	// init cleanup service & schedule task
	chanErr := make(chan error)
	cleanupService := services.NewCleanupUploadsService(dbService)
	if err := schedule.Schedule(10, cleanupService, chanErr); err != nil {
		log.Fatalf("Failed to schedule cleanup task: %v", err)
	}

	go func() {
		for err := range chanErr {
			log.Printf("[Schedule] Task failed: %v", err)
		}
	}()

	// init router & run
	route := router.NewRoute(dbService)
	if err := route.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}

	bootstrap(port, mode)
}
