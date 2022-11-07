package main

import (
	cryptoAPI "cryptoAPI/src/postgressdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := cryptoAPI.ConnectToDB()
	defer db.Close()
	fmt.Println("And we're live bitches ")
	setupRouter()
}

func setupRouter() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome - the database is up",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatal()
	}
}
