package sale

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/TerrexTech/go-kafkautils/kafka"

	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/TerrexTech/uuuid"
	"github.com/pkg/errors"
)

var producer *kafka.Producer

// Insert handles "insert" events.
func Insert(collection *mongo.Collection, event *model.Event) *model.KafkaResponse {
	switch event.ServiceAction {
	case "saleValidated":
		return saleValidated(collection, event)
	default:
		return saleCreated(collection, event)
	}
}

func updateSale(s *Sale) bool {
	if len(s.Items) == 0 {
		return false
	}
	marshalItems, err := json.Marshal(s)
	if err != nil {
		err = errors.Wrap(err, "Error marshalling SaleItems")
		log.Println(err)
		return false
	}

	uid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "Error generating UUID")
		log.Println(err)
		return false
	}
	e := model.Event{
		AggregateID:   2,
		EventAction:   "update",
		ServiceAction: "createSale",
		Data:          marshalItems,
		NanoTime:      time.Now().UnixNano(),
		UUID:          uid,
		Version:       0,
		YearBucket:    2018,
	}

	if producer == nil {
		kafkaBrokersStr := os.Getenv("KAFKA_BROKERS")
		producer, err = kafka.NewProducer(&kafka.ProducerConfig{
			KafkaBrokers: *commonutil.ParseHosts(kafkaBrokersStr),
		})
		if err != nil {
			err = errors.Wrap(err, "ValidateSale: Error creating producer")
			log.Println(err)
			return false
		}
	}

	topic := os.Getenv("KAFKA_PRODUCER_EVENT_TOPIC")
	result, err := json.Marshal(e)
	if err != nil {
		err = errors.Wrap(err, "Insert: Error marshalling result")
		log.Println(err)
		return false
	}
	producer.Input() <- kafka.CreateMessage(topic, result)
	return true
}
