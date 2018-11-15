package mockutil

import (
	"encoding/json"
	"log"
	"time"

	"github.com/TerrexTech/go-kafkautils/kafka"

	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
)

var producer *kafka.Producer
var topic = "event.rns_eventstore.events"

func GenEvent(
	aggregateID int8,
	data interface{},
	eventAction string,
) (*model.Event, error) {
	marshalData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	uuid, err := uuuid.NewV4()
	if err != nil {
		return nil, err
	}
	cid, err := uuuid.NewV4()
	if err != nil {
		return nil, err
	}

	event := &model.Event{
		AggregateID:   aggregateID,
		CorrelationID: cid,
		EventAction:   eventAction,
		Data:          marshalData,
		NanoTime:      time.Now().UnixNano(),
		UUID:          uuid,
		YearBucket:    2018,
	}
	log.Println(uuid)

	if producer == nil {
		producer, err = kafka.NewProducer(&kafka.ProducerConfig{
			KafkaBrokers: []string{"kafka:9092"},
		})
	}

	marshalEvent, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	msg := kafka.CreateMessage(topic, marshalEvent)
	producer.Input() <- msg
	return event, nil
}
