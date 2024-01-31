package sender

import (
	"encoding/json"
	"sender_wb/internal/random"
	"time"

	"github.com/nats-io/stan.go"
)

func SendMessage(clusterID, clientID string) {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		panic("can't connect " + err.Error())
	}
	defer sc.Close()

	for {
		data, _ := json.Marshal(random.RandomNormal())
		err = sc.Publish("orderUpdates", data)
		if err != nil {
			panic("can't send message")
		}

		time.Sleep(time.Second)
	}
}
