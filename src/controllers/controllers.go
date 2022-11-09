package controllers

import (
	cryptoAPI "cryptoAPI/src/postgressdb"
	"github.com/gin-gonic/gin"
	"log"
)

type Crypto struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Amount_Owned int    `json:"amount_owned"`
	Image_Src    string `json:"image_src"`
}

func GetCryptos(c *gin.Context) {

	db := cryptoAPI.ConnectToDB()

	defer db.Close()
	results, err := db.Query("SELECT * FROM crypto")
	if err != nil {

		log.Fatal(err)
	}
	cryptos := []Crypto{}
	for results.Next() {
		var crypto Crypto

		err = results.Scan(&crypto.ID, &crypto.Name, &crypto.Amount_Owned, &crypto.Image_Src)
		if err != nil {

			log.Fatal(err)
		}
		cryptos = append(cryptos, crypto)
	}
	c.JSON(200, cryptos)
	return
}
