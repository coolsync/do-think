package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	bj38 "bj38/proto/bj38"
)

type Bj38 struct{}

func (e *Bj38) Handle(ctx context.Context, msg *bj38.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *bj38.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
