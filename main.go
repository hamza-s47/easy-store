package main

import (
	// "fmt"

	"github.com/hamza-s47/easy-store/config"
	// "github.com/hamza-s47/easy-store/routes"
)

func main() {
	env := config.LoadEnv()

	// r := routes.Routes_init()
	// r.Run(":8080")

	config.InitDB(env)
}
