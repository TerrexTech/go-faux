package mockutil

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TerrexTech/uuuid"
	"github.com/pkg/errors"
)

func GenUUID() uuuid.UUID {
	uuid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "Error generating UUID")
		return uuuid.UUID{}
	}
	return uuid
}

func GenInt(min int, max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(max-min) + min
}

func GenFloat(min float64, max float64) float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	random := min + r1.Float64()*(max-min)
	return random
}

func GenString(chars string, length int) string {
	if chars == "" {
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	}
	charLength := len(chars)

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)

	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		err = errors.Wrap(err, "Error while generating random String")
		log.Println(err)
		return ""
	}
	for i, b := range bytes {
		bytes[i] = chars[b%byte(charLength)]
	}

	return string(bytes)
}

func GenBarcode() string {
	chars := "0123456789"
	barcode := GenString(chars, 12)
	return barcode
}

func GenSKU() string {
	s1 := GenString("", 3)
	s2 := GenString("", 3)
	s3 := GenString("", 3)
	s4 := GenString("", 2)

	return fmt.Sprintf("%s-%s-%s-%s", s1, s2, s3, s4)
}

func AvgWeight(fruit string) float64 {
	dictionary := map[string]float64{
		"Apple":        0.33,
		"Banana":       0.26,
		"Grapes":       0.01,
		"Lettuce":      0.028,
		"Mango":        0.44,
		"Orange":       0.30,
		"Pear":         0.44,
		"Strawberry":   0.026,
		"Sweet Pepper": 0.992,
		"Tomato":       0.328,
	}

	if dictionary[fruit] == 0 {
		return GenFloat(0, 1)
	}
	return dictionary[fruit]
}

func GenFruitName() string {
	dictionary := []string{
		"Apple",
		"Banana",
		"Grapes",
		"Lettuce",
		"Mango",
		"Orange",
		"Pear",
		"Strawberry",
		"Sweet Pepper",
		"Tomato",
	}

	index := GenInt(0, len(dictionary))
	return dictionary[index]
}

func GenOrigin() string {
	dictionary := []string{
		"ON Canada",
		"BC Canada",
		"SK Canada",
		"MN Canada",
		"NS Canada",
		"PEI Canada",
		"QC Canada",
	}

	index := GenInt(0, len(dictionary))
	return dictionary[index]
}

func GenLot() string {
	s1 := GenString("", 2)
	i1 := GenInt(0, 9999)

	return fmt.Sprintf("%s%d", s1, i1)
}
