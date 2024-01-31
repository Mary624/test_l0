package event

import (
	"encoding/json"
	"log/slog"
	"test-go/internal/logger"
	"test-go/internal/storage"

	"github.com/go-playground/validator"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
)

type Saver interface {
	SaveOrder(storage.Order) error
}

type EventHandler struct {
	sc stan.Conn
}

func New(clusterID, clientID string) (EventHandler, error) {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		return EventHandler{}, err
	}
	return EventHandler{
		sc: sc,
	}, nil
}

func (e EventHandler) HandleEvent(log *slog.Logger, subjectNats, durableName string, saver Saver, c *cache.Cache) error {

	_, err := e.sc.Subscribe(subjectNats, func(msg *stan.Msg) {
		var order storage.Order
		json.Unmarshal(msg.Data, &order)
		SaveOrder(order, log, saver, c)
	}, stan.DurableName(durableName))
	if err != nil {
		return nil
	}

	return nil
}

func SaveOrder(order storage.Order, log *slog.Logger, saver Saver, c *cache.Cache) error {
	if err := validator.New().Struct(order); err != nil {
		log.Error("error: validate", logger.Err(err))
		return err
	}
	c.Set(order.OrderUid, order, cache.DefaultExpiration)

	err := saver.SaveOrder(order)
	if err != nil {
		log.Error("can't save order", logger.Err(err))
		return err
	}
	return nil
}
