package cryptoAPI

import (
	"database/sql"
	"flag"
	"fmt"
)

type DB struct {
	DB *sql.DB
}

func ConnectToDB() *sql.DB {
	dbUrl := flag.String("url", "no flag provided", "dburl")
	dbPass := flag.String("password", "No password provided", "dbpass")
	dbUsername := flag.String("username", "no username provided", "dbuser")

	flag.Parse()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		*dbUrl, 5432, *dbUsername, *dbPass, *dbUsername)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
