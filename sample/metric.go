package sample

import (
	"log"
	"time"

	"github.com/TerrexTech/go-faux/model"
	"github.com/pkg/errors"

	"github.com/TerrexTech/go-mongoutils/mongo"

	"github.com/TerrexTech/go-faux/mockutil"
)

func Metric(deviceColl *mongo.Collection) (*model.Metric, error) {
	device, err := mockutil.RandomDevice(deviceColl)
	if err != nil {
		err = errors.Wrap(err, "Error getting random-device")
		return nil, err
	}

	co2 := mockutil.GenFloat(300, 1200)

	log.Printf("-33333333333333333")
	log.Printf("%+v", device)

	return &model.Metric{
		MetricID:      mockutil.GenUUID(),
		ItemID:        device.ItemID,
		DeviceID:      device.DeviceID,
		Timestamp:     time.Now().Unix(),
		TempIn:        mockutil.GenFloat(21, 27),
		Humidity:      mockutil.GenFloat(40, 80),
		CarbonDioxide: co2,
		Ethylene:      co2 / 400,
		SKU:           device.SKU,
		Name:          device.Name,
		Lot:           device.Lot,
	}, nil
}
