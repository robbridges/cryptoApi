package main

import (
	router "cryptoAPI/src/Router"
	cryptoAPI "cryptoAPI/src/postgressdb"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	_ = flag.String("url", "no flag provided", "dburl")
	_ = flag.String("password", "No password provided", "dbpass")
	_ = flag.String("username", "no username provided", "dbuser")

	flag.Parse()
	db := cryptoAPI.ConnectToDB()
	defer db.Close()
	fmt.Println("And we're live bitches ")
	fmt.Println(flag.Lookup("url").Value.(flag.Getter).Get())
	router.SetupRouter()
}
