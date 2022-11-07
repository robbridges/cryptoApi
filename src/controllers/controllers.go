package controllers

import cryptoAPI "cryptoAPI/src/postgressdb"

func InsertIntoCryptoTable(cryptoName string, cryptoAmount int, cryptoImageAddress string) {
	db := cryptoAPI.ConnectToDB()
	sqlStatment := `
	INSERT INTO crypto (name, amount_owned, image_src)
	VALUES($1, $2, $3)`
	_, err := db.Exec(sqlStatment, cryptoName, cryptoAmount, cryptoImageAddress)
	if err != nil {
		panic(err)
	}
	
}
