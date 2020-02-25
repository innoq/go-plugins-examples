package utils

import (
	"math/rand"
	"github.com/lithammer/shortuuid"
)

type DataGenerator struct {}

func NewDataGenerator() *DataGenerator {
	return &DataGenerator{}
}

func (d *DataGenerator) Next() (map[string]string) {
	data := make(map[string]string)
	i := rand.Intn(100)
	if i > 66 {
		data["type"] = "user"
	} else {
		data["type"] = "object"
	}
	data["id"] = shortuuid.New()
	return data
}
