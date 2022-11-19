package Router

import (
	"cryptoAPI/src/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome - the database is up",
		})
	})
	router.GET("/crypto", controllers.GetCryptos)
	router.GET("/crypto/:id", controllers.GetCryptoById)
	router.POST("/crypto/create", controllers.CreateCrypto)
	router.DELETE("/crypto/delete/:id", controllers.DeleteCoin)
	router.PATCH("/crypto/update/amount/:id", controllers.UpdateCoinAmountOwned)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal()
	}
}
