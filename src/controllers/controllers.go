package controllers

import (
	cryptoAPI "cryptoAPI/src/postgressdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func GetCryptoById(c *gin.Context) {
	db := cryptoAPI.ConnectToDB()

	coin := Crypto{
		-1,
		"not found",
		0,
		"not found",
	}

	id := c.Param("id")
	sqlStatement := `SELECT id, name, amount_owned, image_src FROM crypto WHERE id = $1;`
	crypto := db.QueryRow(sqlStatement, id)
	fmt.Println(id)
	err := crypto.Scan(&coin.ID, &coin.Name, &coin.Amount_Owned, &coin.Image_Src)
	if err != nil {
		c.JSON(400, "Error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, coin)
	return
}
