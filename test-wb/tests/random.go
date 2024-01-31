package tests

import (
	"math/rand"
	"strconv"
	"strings"
	"test-go/internal/storage"
	"time"

	"github.com/brianvoe/gofakeit"
)

func RandomNormal() storage.Order {
	gofakeit.Seed(time.Now().UnixMilli())
	return storage.Order{
		OrderUid:    gofakeit.UUID(),
		TrackNumber: RandomStr(),
		Entry:       RandomStr(),
		Delivery: storage.Delivery{
			Name:    gofakeit.Name(),
			Phone:   gofakeit.Phone(),
			Zip:     strconv.Itoa(int(gofakeit.Uint32())),
			City:    gofakeit.City(),
			Address: gofakeit.Address().Address,
			Region:  RandomStr(),
			Email:   gofakeit.Email(),
		},
		Payment: storage.Payment{
			Transaction:  gofakeit.UUID(),
			RequestId:    RandomStr(),
			Currency:     RandomStr(),
			Provider:     RandomStr(),
			Amount:       int64(gofakeit.Uint32()),
			PaymentDt:    int64(gofakeit.Uint32()),
			Bank:         RandomStr(),
			DeliveryCost: int64(gofakeit.Uint32()),
			GoodsTotal:   int64(gofakeit.Uint32()),
			CustomFee:    int64(gofakeit.Uint32()),
		},
		Items:             RandomNormalItems(),
		Locale:            RandomStr(),
		InternalSignature: RandomStr(),
		CustomerId:        RandomStr(),
		DeliveryService:   RandomStr(),
		Shardkey:          strconv.Itoa(int(gofakeit.Uint32())),
		SmId:              int64(gofakeit.Uint32()),
		DateCreated:       time.Now(),
		OofShard:          strconv.Itoa(int(gofakeit.Uint32())),
	}
}

func RandomNormalItems() []storage.Item {
	min, max := 1, 20
	c := rand.Intn(max-min) + min
	items := make([]storage.Item, 0, c)
	for i := 0; i < c; i++ {
		items = append(items, storage.Item{
			ChrtId:      int64(gofakeit.Uint32()),
			TrackNumber: RandomStr(),
			Price:       int64(gofakeit.Uint32()),
			Rid:         gofakeit.UUID(),
			Name:        gofakeit.Word(),
			Sale:        int64(gofakeit.Uint32()),
			Size:        RandomStr(),
			TotalPrice:  int64(gofakeit.Uint32()),
			NmId:        int64(gofakeit.Uint32()),
			Brand:       RandomStr(),
			Status:      int64(gofakeit.Uint32()),
		})
	}
	return items
}

func RandomStr() string {
	var b strings.Builder
	min, max := 1, 20
	c := rand.Intn(max-min) + min
	chs := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < c; i++ {
		ch := ([]rune)(chs)[rand.Intn(len(chs)-0)+0]
		b.WriteRune(ch)
	}
	return b.String()
}

func RandomWithoutItems() storage.Order {
	order := RandomNormal()
	order.Items = make([]storage.Item, 0, 0)
	return order
}

func RandomtrackNumber() storage.Order {
	order := RandomNormal()
	order.TrackNumber = ""
	return order
}
