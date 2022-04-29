package api

import (
	"fmt"
	"log"
	"os"

	"github.com/INI-MED/gocrud/api/controllers"
	"github.com/INI-MED/gocrud/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("POSTGRES_DB"),
	)

	seed.Load(server.DB)

	server.Run(":8080")
}
