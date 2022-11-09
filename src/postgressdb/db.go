package cryptoAPI

import (
	"database/sql"
	"flag"
	"fmt"
)

func ConnectToDB() *sql.DB {
	dbUrl := flag.Lookup("url").Value.(flag.Getter).Get()
	dbPass := flag.Lookup("password").Value.(flag.Getter).Get()
	dbUsername := flag.Lookup("username").Value.(flag.Getter).Get()

	flag.Parse()
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
