package main

import (
	"context"
	"log"
	"os"

	"PaginationPlayground/internal/client"
	"PaginationPlayground/internal/handler"
	"PaginationPlayground/internal/persist"
	"PaginationPlayground/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbContext := persist.NewDatabaseContext(ctx, os.Getenv("DB_URL"))
	itemRepo := persist.NewItemRepository(dbContext)
	itemClient := client.NewOsrsClient()
	itemService := service.NewOsrsService(itemRepo, itemClient)
	itemHandler := handler.NewOsrsHandler(itemService)
	r := gin.Default()

	r.GET("/fetch-osrs-data", itemHandler.FetchAndPersistItems)

	r.GET("/search-item/:name", itemHandler.SearchItems)

	_ = r.Run(":8080")
}
