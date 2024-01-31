package tests

import (
	"log"
	"os"
	"path/filepath"
	"test-go/internal/config"
	"test-go/internal/event"
	discaredlogger "test-go/internal/logger/discaredLogger"
	"test-go/internal/storage"
	"test-go/internal/storage/postgres"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullPath := filepath.Join(filepath.Join(path, ".."), ".env")
	err = godotenv.Load(fullPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	type v struct {
		Order storage.Order
		Err   bool
	}
	l := []v{
		{
			Order: RandomNormal(),
			Err:   false,
		},
		{
			Order: storage.Order{},
			Err:   true,
		},
		{
			Order: RandomWithoutItems(),
			Err:   true,
		},
		{
			Order: RandomtrackNumber(),
			Err:   true,
		},
	}
	cfg := config.MustLoad()
	db, err := postgres.New(cfg.DBConfig)
	if err != nil {
		panic("can't connect to db")
	}
	log := discaredlogger.NewDiscardLogger()
	c := cache.New(5*time.Minute, 10*time.Minute)
	for _, value := range l {
		err := event.SaveOrder(value.Order, log, db, c)
		b := false
		if err != nil {
			b = true
		}
		assert.Equal(t, b, value.Err)
	}
}
