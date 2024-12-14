package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"stock-sleuth-backend/internal/config"
	"stock-sleuth-backend/internal/handlers"
	"stock-sleuth-backend/internal/services"
	"stock-sleuth-backend/pkg/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	apiClient := api.NewHTTPClient()

	stockService := services.NewStockService(cfg.FmpAPIKey, apiClient)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	stockHandler := handlers.NewStockHandler(stockService)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/stocks/:symbol", stockHandler.GetStockInfo)
	}

	serverAddr := cfg.ServerAddress()
	log.Println("Listening on " + serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
