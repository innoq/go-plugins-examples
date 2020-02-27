package utils

import (
	"math/rand"

	"github.com/lithammer/shortuuid"
)

// DataGenerator - Tool for creating random test data
type DataGenerator struct{}

// NewDataGenerator - create a new data generator
func NewDataGenerator() *DataGenerator {
	return &DataGenerator{}
}

// Next - generate the next data entry
func (d *DataGenerator) Next() map[string]string {
	data := make(map[string]string)
	i := rand.Intn(100)
	if i > 66 {
		data["type"] = "user"
	} else {
		data["type"] = "object"
	}
	if i > 75 {
		data["event"] = "create"
	} else if i > 50 {
		data["event"] = "update"
	} else if i > 25 {
		data["event"] = "delete"
	} else {
		data["event"] = "read"
	}
	data["id"] = shortuuid.New()
	return data
}
