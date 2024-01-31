package main

import (
	"log"
	"test-go/internal/config"
	"test-go/internal/server"
	"test-go/internal/storage/postgres"

	"github.com/joho/godotenv"

	"os"
	"path/filepath"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullPath := filepath.Join(filepath.Join(path, "../.."), ".env")
	err = godotenv.Load(fullPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	cfg := config.MustLoad()

	db, err := postgres.New(cfg.DBConfig)
	if err != nil {
		panic("can't connect to db")
	}

	s := server.New(cfg, db, db)
	err = s.Run(cfg.Port)
	if err != nil {
		panic(err)
	}
}
