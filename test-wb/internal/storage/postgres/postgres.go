package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"test-go/internal/config"
	"test-go/internal/storage"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(cfg config.DBConfig) (*Storage, error) {
	const op = "storage.New"

	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.HostDB, cfg.PortDB, cfg.UserDB, cfg.PassDB, cfg.DBName)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveOrder(order storage.Order) error {
	const op = "storage.SaveOrder"

	stmt, err := s.db.Prepare(`INSERT INTO orders(order_uid, info)
	 VALUES($1, $2);`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = stmt.Exec(order.OrderUid, order)
	if pgerr, ok := err.(*pq.Error); ok {
		if pgerr.Code == "23505" {
			return fmt.Errorf("%s: %w", op, storage.ErrEntryExists)
		}
	}
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetOrderById(id string) (storage.Order, error) {
	const op = "storage.GetOrderById"

	stmt, err := s.db.Prepare(`SELECT info FROM orders WHERE order_uid=$1;`)
	if err != nil {
		return storage.Order{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRow(id)
	var order storage.Order
	err = row.Scan(&order)
	if errors.Is(sql.ErrNoRows, err) {
		return storage.Order{}, fmt.Errorf("%s: %w", op, storage.ErrEntryNotFound)
	}
	if err != nil {
		return storage.Order{}, fmt.Errorf("%s: %w", op, err)
	}
	return order, nil
}

func (s *Storage) GetOrdes(limit int) (*cache.Cache, error) {
	const op = "storage.GetOrderById"

	stmt, err := s.db.Prepare(`SELECT info FROM orders LIMIT $1;`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	res := cache.New(5*time.Minute, 10*time.Minute)
	for rows.Next() {
		var order storage.Order
		err = rows.Scan(&order)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		res.Set(order.OrderUid, order, cache.DefaultExpiration)
	}
	return res, nil
}
