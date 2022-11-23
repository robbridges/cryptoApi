package main

import (
	router "cryptoAPI/src/Router"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	ginRouter := router.SetupRouter()
	err := ginRouter.Run(":8080")
	if err != nil {
		log.Fatal()
	}

}

func test() {
	fmt.Println("test")
}
