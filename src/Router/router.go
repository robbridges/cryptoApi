package Router

import (
	"cryptoAPI/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome - the database is up",
		})
	})
	cryptoRoute := router.Group("/crypto")
	{
		cryptoRoute.GET("/allcrypto", controllers.GetCryptos)
		cryptoRoute.GET("/:id", controllers.GetCryptoById)
		cryptoRoute.POST("/create", controllers.CreateCrypto)
		cryptoRoute.DELETE("/delete/:id", controllers.DeleteCoin)
		cryptoRoute.PATCH("update/amount/:id", controllers.UpdateCoinAmountOwned)
	}

	return router
}
