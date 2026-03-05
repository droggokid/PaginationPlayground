package main

import (
	"PaginationPlayground/internal"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	h := internal.NewHttpClient()
	r := gin.Default()

	r.GET("/get-items", func(c *gin.Context) {
		err := h.GetOSRSData(c)
		if err != nil {
			fmt.Println(err)
		}
	})

	_ = r.Run(":8080")
}
