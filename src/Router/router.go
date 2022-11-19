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
	cryptoRoute := router.Group("/crypto")
	{
		cryptoRoute.GET("/", controllers.GetCryptos)
		cryptoRoute.GET("/:id", controllers.GetCryptoById)
		cryptoRoute.POST("/create", controllers.CreateCrypto)
		cryptoRoute.DELETE("/delete/:id", controllers.DeleteCoin)
		cryptoRoute.PATCH("update/amount/:id", controllers.UpdateCoinAmountOwned)
	}

	err := router.Run(":8080")
	if err != nil {
		log.Fatal()
	}
}
