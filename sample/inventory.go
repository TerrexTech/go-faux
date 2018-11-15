package sample

import (
	"time"

	"github.com/TerrexTech/go-faux/mockutil"

	"github.com/TerrexTech/agg-inventory-cmd/inventory"
)

func Inventory() *inventory.Inventory {
	return &inventory.Inventory{
		ItemID:       mockutil.GenUUID(),
		Barcode:      mockutil.GenBarcode(),
		DateArrived:  time.Now().Unix(),
		DeviceID:     mockutil.GenUUID(),
		Lot:          mockutil.GenLot(),
		Name:         mockutil.GenFruitName(),
		Origin:       mockutil.GenOrigin(),
		Price:        mockutil.GenFloat(1, 10),
		RSCustomerID: mockutil.GenUUID(),
		SalePrice:    mockutil.GenFloat(1, 6),
		SKU:          mockutil.GenSKU(),
		Timestamp:    time.Now().Unix(),
		TotalWeight:  mockutil.GenFloat(300, 1000),
		UPC:          int64(mockutil.GenInt(10, 100)),
	}
}
