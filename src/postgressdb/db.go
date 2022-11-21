package cryptoAPI

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func ConnectToDB() *sql.DB {
	viper.SetConfigFile("variables.env")
	viper.ReadInConfig()
	dbUrl := viper.Get("URL")
	dbPass := viper.Get("PASSWORD")
	dbUsername := viper.Get("USERNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		dbUrl, 5432, dbUsername, dbPass, dbUsername)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
