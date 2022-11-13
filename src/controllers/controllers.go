package controllers

import (
	coin "cryptoAPI/src/models"
	cryptoAPI "cryptoAPI/src/postgressdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Crypto coin.Crypto

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

			c.JSON(http.StatusInternalServerError, "Error :"+err.Error())
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
		c.JSON(http.StatusNotFound, "Error: "+err.Error())
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
	coin.Amount_Owned, err = strconv.ParseFloat(cryptoAmount, 64)
	coin.Image_Src = cryptoImage

	c.JSON(201, coin)
	return
}

func DeleteCoin(c *gin.Context) {
	db := cryptoAPI.ConnectToDB()
	defer db.Close()
	id := c.Param("id")

	var coin Crypto
	sqlStatement := `SELECT id FROM crypto WHERE id = $1;`
	crypto := db.QueryRow(sqlStatement, id)
	err := crypto.Scan(&coin.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, "Error: "+err.Error())
		return
	}

	sqlStatement = `
	DELETE FROM crypto
	WHERE id = $1;`

	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Error: "+err.Error())
	}
	c.Status(http.StatusOK)
	return

}

func UpdateCoinAmountOwned(c *gin.Context) {
	db := cryptoAPI.ConnectToDB()
	defer db.Close()

	id := c.Param("id")

	if c.Query("amount_owned") == "" {
		c.JSON(http.StatusBadRequest, "An amount_owned query string is required")
		return
	}
	amountOwned := c.Query("amount_owned")

	sqlStatement := `
	UPDATE crypto
	SET amount_owned = $2
	WHERE id = $1;`

	_, err := db.Exec(sqlStatement, id, amountOwned)

	if err != nil {
		fmt.Println("firstError " + err.Error())
		c.Status(500)
		return
	}
	var updatedCoin Crypto
	sqlStatement = `SELECT id, name, amount_owned, image_src FROM crypto WHERE id = $1;`
	crypto := db.QueryRow(sqlStatement, id)
	err = crypto.Scan(&updatedCoin.ID, &updatedCoin.Name, &updatedCoin.Amount_Owned, &updatedCoin.Image_Src)
	if err != nil {
		fmt.Println("Second Error " + err.Error())
		c.Status(500)
		return
	}

	c.JSON(http.StatusOK, updatedCoin)
}
