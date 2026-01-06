package main

import (
	// "fmt"

	"log"

	"github.com/hamza-s47/easy-store/config"
	// "github.com/hamza-s47/easy-store/routes"
)

func main() {
	env := config.LoadEnv()

	// Gin
	// r := routes.Routes_init()
	// r.Run(":8080")

	//DB
	db, err := config.InitDB(env)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
