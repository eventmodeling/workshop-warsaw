package infrastructure

import (
	"encoding/json"
	"io"
	"log"
)

type EventPublisher struct {
	writer io.WriteCloser
}

func NewEventPublisher(writer io.WriteCloser) EventPublisher {
	if writer == nil {
		panic("writer is nil")
	}

	return EventPublisher{
		writer: writer,
	}
}

func (p EventPublisher) Publish(event interface{}) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	log.Print("Publishing event:", string(payload))

	_, err = p.writer.Write(payload)
	return err
}

func (p EventPublisher) Close() error {
	return p.writer.Close()
}
