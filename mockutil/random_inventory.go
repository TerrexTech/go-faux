package mockutil

import (
	"log"

	"github.com/TerrexTech/go-faux/model"

	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/pkg/errors"
)

func RandomInventory(coll *mongo.Collection) (*model.Inventory, error) {
	params := map[string]interface{}{
		"timestamp": map[string]interface{}{
			"$ne": 0,
		},
	}

	findopts := findopt.Limit(100)
	results, err := coll.Find(params, findopts)
	if err != nil {
		err = errors.Wrap(err, "RandomInventory: Error in Find")
		log.Println(err)
		return nil, err
	}

	randomIndex := GenInt(1, len(results)-2)
	inv, assertOK := results[randomIndex].(*model.Inventory)
	if !assertOK {
		err = errors.New("RandomInventory: error asserting inventory")
		log.Println(err)
		return nil, err
	}

	return inv, nil
}
