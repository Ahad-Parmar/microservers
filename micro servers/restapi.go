  
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	a := gin.Default()
	a.GET("/", func(b *gin.Context) {
		b.JSON(200, gin.H{
			"message": "Heyy World!!",
		})
	})
	a.Run()
}