package mockutil

import (
	"log"

	"github.com/TerrexTech/go-faux/model"

	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/pkg/errors"
)

func RandomDevice(coll *mongo.Collection) (*model.Device, error) {
	params := map[string]interface{}{
		"timestamp": map[string]interface{}{
			"$ne": 0,
		},
	}

	findopts := findopt.Limit(100)
	results, err := coll.Find(params, findopts)
	if err != nil {
		err = errors.Wrap(err, "RandomDevice: Error in Find")
		log.Println(err)
		return nil, err
	}

	log.Printf("%+v", results)
	randomIndex := GenInt(1, len(results)-2)
	log.Printf("%+v", results[randomIndex])
	log.Printf("/////////////////")
	device, assertOK := results[randomIndex].(*model.Device)
	if !assertOK {
		err = errors.New("RandomDevice: error asserting device")
		log.Println(err)
		return nil, err
	}

	return device, nil
}
