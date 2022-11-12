package main

import (
	controllers "cryptoAPI/src/controllers"
	cryptoAPI "cryptoAPI/src/postgressdb"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	_ = flag.String("url", "no flag provided", "dburl")
	_ = flag.String("password", "No password provided", "dbpass")
	_ = flag.String("username", "no username provided", "dbuser")

	flag.Parse()
	db := cryptoAPI.ConnectToDB()
	defer db.Close()
	fmt.Println("And we're live bitches ")
	fmt.Println(flag.Lookup("url").Value.(flag.Getter).Get())
	setupRouter()
}

func setupRouter() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome - the database is up",
		})
	})
	router.GET("/crypto", controllers.GetCryptos)
	router.GET("/crypto/:id", controllers.GetCryptoById)
	router.POST("/crypto/create", controllers.CreateCrypto)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		log.Fatal()

	}
}
