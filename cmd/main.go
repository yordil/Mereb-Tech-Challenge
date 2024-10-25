package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yordil/mereb-tech-challenge/internal/route"

	"github.com/yordil/mereb-tech-challenge/docs"
)


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// allowed to accept  a request from every origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		// Handle OPTIONS requests for preflight checks
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func main() {
	
	router := gin.Default()
	// Add CORS middleware
	router.Use(CORSMiddleware())
	
	fmt.Println(docs.SwaggerInfo.Title)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	route.SetupAllRoutes(router)

	// Configure the HTTP server with the router and start the server
	server := &http.Server{
		Addr:    ":8080",       
		Handler: router,
	}



	// Start server in a goroutine to handle graceful shutdown and prevent memory leak
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
