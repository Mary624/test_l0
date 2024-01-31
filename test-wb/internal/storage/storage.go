package storage

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrEntryNotFound = errors.New("order not found")
	ErrEntryExists   = errors.New("order exists")
)

type Order struct {
	OrderUid          string    `json:"order_uid" validate:"required"`
	TrackNumber       string    `json:"track_number" validate:"required"`
	Entry             string    `json:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	SmId              int64     `json:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}

type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int64  `json:"amount" validate:"required"`
	PaymentDt    int64  `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int64  `json:"delivery_cost" validate:"required"`
	GoodsTotal   int64  `json:"goods_total" validate:"required"`
	CustomFee    int64  `json:"custom_fee"`
}

type Item struct {
	ChrtId      int64  `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int64  `json:"sale" validate:"required"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  int64  `json:"total_price" validate:"required"`
	NmId        int64  `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      int64  `json:"status" validate:"required"`
}

func (o Order) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *Order) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}
