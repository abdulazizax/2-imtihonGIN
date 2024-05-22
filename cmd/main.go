package main

import (
	"exam/api"
	"exam/config"
	"exam/storage"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDB(path string) (*sqlx.DB, error) {
	cfg := config.Load(path)

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.DbHost,
		cfg.Postgres.DbPort,
		cfg.Postgres.DbUser,
		cfg.Postgres.DbPassword,
		cfg.Postgres.DbName,
	)

	db, err := sqlx.Connect("postgres", psqlUrl)
	return db, err
}

func main() {
	path := ".env"
	db, err := GetDB(path)
	if err != nil {
		log.Fatal("Failed connecting to database: ", err)
	}

	storage := storage.NewStoragePg(db)

	server := api.New(api.Option{
		Storage: storage,
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	if err := server.Run(":" + port); err != nil {
		log.Fatal("Failed to run HTTP server: ", err)
	}

	log.Print("Server stopped")
}
