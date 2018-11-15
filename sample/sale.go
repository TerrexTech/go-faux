package sample

import (
	"log"
	"time"

	"github.com/TerrexTech/go-faux/model"

	"github.com/pkg/errors"

	"github.com/TerrexTech/go-mongoutils/mongo"

	"github.com/TerrexTech/go-faux/mockutil"
)

func Sale(invColl *mongo.Collection) (*model.Sale, error) {
	items := []model.SoldItem{}

	numItems := mockutil.GenInt(1, 10)
	for i := 0; i < numItems; i++ {
		inv, err := mockutil.RandomInventory(invColl)
		if err != nil {
			err = errors.Wrap(err, "Error generating random inventory")
			log.Println(err)
			continue
		}
		soldItem := model.SoldItem{
			ItemID: inv.ItemID,
			UPC:    inv.UPC,
			Weight: mockutil.GenFloat(1, inv.TotalWeight-1),
			Lot:    inv.Lot,
			SKU:    inv.SKU,
		}
		items = append(items, soldItem)
	}

	return &model.Sale{
		SaleID:    mockutil.GenUUID(),
		Items:     items,
		Timestamp: time.Now().Unix(),
	}, nil
}
