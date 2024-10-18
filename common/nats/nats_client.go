package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

var NC *nats.Conn

func InitNATS() {
    var err error
    NC, err = nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatal(err)
    }
}

func Publish(subject string, data []byte) error {
    return NC.Publish(subject, data)
}

func Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
    return NC.Subscribe(subject, handler)
}