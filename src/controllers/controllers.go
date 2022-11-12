package controllers

import (
	cryptoAPI "cryptoAPI/src/postgressdb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	defer db.Close()
	var coin Crypto

	id := c.Param("id")
	sqlStatement := `SELECT id, name, amount_owned, image_src FROM crypto WHERE id = $1;`
	crypto := db.QueryRow(sqlStatement, id)
	err := crypto.Scan(&coin.ID, &coin.Name, &coin.Amount_Owned, &coin.Image_Src)
	if err != nil {
		c.JSON(400, "Error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, coin)
	return
}

func CreateCrypto(c *gin.Context) {
	db := cryptoAPI.ConnectToDB()
	defer db.Close()

	cryptoName := c.Query("name")
	cryptoAmount := c.Query("amount_owned")
	cryptoImage := c.Query("image_src")

	var coin Crypto
	id := 0

	sqlStatement := `INSERT INTO crypto (name, amount_owned, image_src)
	VALUES($1, $2, $3)
	RETURNING id`

	err := db.QueryRow(sqlStatement, cryptoName, cryptoAmount, cryptoImage).Scan(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	coin.ID = id
	coin.Name = cryptoName
	coin.Amount_Owned, err = strconv.Atoi(cryptoAmount)
	coin.Image_Src = cryptoImage

	c.JSON(201, coin)
	return
}

func deleteCoin(c *gin.Context) {
	db := cryptoAPI.ConnectToDB()
	defer db.Close()

}
