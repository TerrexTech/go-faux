package main

import (
	"log"
	"time"

	"github.com/pkg/errors"

	"github.com/TerrexTech/go-faux/mockutil"

	"github.com/TerrexTech/go-faux/sample"

	"github.com/TerrexTech/go-faux/collection"
	"github.com/TerrexTech/go-mongoutils/mongo"
)

func main() {
	config := mongo.ClientConfig{
		Hosts:               []string{"localhost:27017"},
		Username:            "root",
		Password:            "root",
		TimeoutMilliseconds: 3000,
	}

	// ====> MongoDB Client
	client, err := mongo.NewClient(config)
	log.Println(err)

	invColl, err := collection.Inventory(client)
	deviceColl, err := collection.Device(client)
	log.Println(err)

	device, err := sample.Device(invColl)
	inventory := sample.Inventory()
	metrics, err := sample.Metric(deviceColl)
	sale, err := sample.Sale(deviceColl)

	mockutil.GenEvent(6, device, "insert")
	log.Println(device)
	log.Println(inventory)
	log.Println(metrics)
	log.Println(sale)
	log.Println(")))))))))))))))")

	go func() {
		deviceCron := mockutil.GenInt(1, 500)
		<-time.After(time.Duration(deviceCron) * time.Second)
		for {
			device, err := sample.Device(invColl)
			if err != nil {
				err = errors.Wrap(err, "Error while getting sample device")
				log.Println(err)
			}
			log.Println("Generating Device...")
			log.Printf("%+v", device)
			mockutil.GenEvent(6, device, "insert")
		}
	}()

	go func() {
		for {
			inventoryCron := mockutil.GenInt(1, 30)
			<-time.After(time.Duration(inventoryCron) * time.Second)
			inventory := sample.Inventory()
			log.Println("Generating Inventory...")
			log.Printf("%+v", inventory)
		}
	}()

	go func() {
		for {
			metricCron := mockutil.GenInt(1, 5)
			<-time.After(time.Duration(metricCron) * time.Second)
			metrics, err := sample.Metric(deviceColl)
			if err != nil {
				err = errors.Wrap(err, "Error while getting sample device")
				log.Println(err)
			}
			log.Println("Generating Metrics...")
			log.Printf("%+v", metrics)
		}
	}()

	for {
		saleCron := mockutil.GenInt(1, 50)
		<-time.After(time.Duration(saleCron) * time.Second)
		sale, err := sample.Sale(deviceColl)
		if err != nil {
			err = errors.Wrap(err, "Error while getting sample device")
			log.Println(err)

		}
		log.Println("Generating Sale...")
		log.Printf("%+v", sale)
	}
}
