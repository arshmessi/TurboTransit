package events

import (
	"TurboTransit/common/nats"
	"encoding/json"
	"log"
)

type EventHandler func(data []byte)

func SubscribeToEvent(subject string, handler EventHandler) {
    _, err := nats.Subscribe(subject, func(msg *nats.Msg) {
        handler(msg.Data)
    })
    if err != nil {
        log.Fatal(err)
    }
}

func PublishEvent(subject string, data interface{}) {
    payload, err := json.Marshal(data)
    if err != nil {
        log.Fatal(err)
    }
    err = nats.Publish(subject, payload)
    if err != nil {
        log.Fatal(err)
    }
}