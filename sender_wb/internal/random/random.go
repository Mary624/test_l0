package random

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

func RandomNormal() Order {
	gofakeit.Seed(time.Now().UnixMilli())
	return Order{
		OrderUid:    gofakeit.UUID(),
		TrackNumber: RandomStr(),
		Entry:       RandomStr(),
		Delivery: Delivery{
			Name:    gofakeit.Name(),
			Phone:   gofakeit.Phone(),
			Zip:     strconv.Itoa(int(gofakeit.Uint32())),
			City:    gofakeit.City(),
			Address: gofakeit.Address().Address,
			Region:  RandomStr(),
			Email:   gofakeit.Email(),
		},
		Payment: Payment{
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

func RandomNormalItems() []Item {
	min, max := 1, 20
	c := rand.Intn(max-min) + min
	items := make([]Item, 0, c)
	for i := 0; i < c; i++ {
		items = append(items, Item{
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
