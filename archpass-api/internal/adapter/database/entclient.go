package database

import (
	"log"
	"os"

	"github.com/garguelles/archpass/ent"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var client *ent.Client

func init() {
	godotenv.Load(`.env`)
	databaseUrl, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatal("Cannot connect to database")
	}

	var err error
	client, err = ent.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalf("Failed opening connection to postgres: %v", err)
	}
}

func EntClient() *ent.Client {
	return client
}
