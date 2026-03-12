package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"PaginationPlayground/internal/client"
	"PaginationPlayground/internal/persist"

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
	client := client.NewOSRSClient(itemRepo)
	r := gin.Default()

	r.GET("/populate-db", func(c *gin.Context) {
		err := client.FetchAndPersistItems(c)
		if err != nil {
			fmt.Println(err)
		}
	})

	r.GET("/search-item/:name", func(c *gin.Context) {
		name := c.Param("name")
		items, err := itemRepo.GetItem(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, items)
	})

	_ = r.Run(":8080")
}
