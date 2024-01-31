package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sender_wb/internal/sender"

	"github.com/joho/godotenv"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullPath := filepath.Join(path, ".env")
	fmt.Println(fullPath)
	err = godotenv.Load(fullPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	clusterID := os.Getenv("CLUSTER_ID")
	if clusterID == "" {
		panic("clusterID is not set")
	}
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		panic("clientID is not set")
	}
	// Раз в секунду отправляет заказ
	sender.SendMessage(clusterID, clientID)
}
