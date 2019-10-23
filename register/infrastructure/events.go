package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	uuid "github.com/satori/go.uuid"
)

type EventPublisher struct {
	EventsDirectory string
}

func (p EventPublisher) Publish(eventName string, event interface{}) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%v_%v", eventName, uuid.NewV4().String())

	file, err := os.Create(path.Join(p.EventsDirectory, filePath))
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	log.Println("Publishing event:", filePath, string(payload))

	_, err = file.Write(payload)
	return err
}
