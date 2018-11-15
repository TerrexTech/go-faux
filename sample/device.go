package sample

import (
	"time"

	"github.com/TerrexTech/go-faux/model"
	"github.com/pkg/errors"

	"github.com/TerrexTech/go-faux/mockutil"
	"github.com/TerrexTech/go-mongoutils/mongo"
)

func Device(invColl *mongo.Collection) (*model.Device, error) {
	inv, err := mockutil.RandomInventory(invColl)
	if err != nil {
		err = errors.Wrap(err, "Error getting random-inventory")
		return nil, err
	}

	return &model.Device{
		ItemID:        inv.ItemID,
		DeviceID:      mockutil.GenUUID(),
		DateInstalled: time.Now().Unix(),
		Lot:           inv.Lot,
		Name:          inv.Name,
		Status:        "Healthy",
		SKU:           inv.SKU,
	}, nil
}
