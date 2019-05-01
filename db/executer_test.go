package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateID(t *testing.T) {

	records := []Record{
		{
			DeviceID:  "mac:123456",
			BirthDate: time.Now().Unix(),
		},
		{
		},
		{
			DeviceID:  "213456789",
			BirthDate: time.Now().Add(time.Hour).Unix(),
		},
	}

	for _, record := range records {
		t.Run(record.DeviceID, func(t *testing.T) {
			assert := assert.New(t)
			assert.NotPanics(func() {
				id := generateID(record)
				assert.NotEmpty(id)
			})

		})

	}
}

func TestConflict(t *testing.T) {
	assert := assert.New(t)

	recordA := Record{
		DeviceID:  "mac:123456",
		BirthDate: time.Now().Unix(),
	}
	recordB := Record{
		DeviceID:  "mac:123456",
		BirthDate: time.Now().Unix(),
	}

	idA := generateID(recordA)
	idB := generateID(recordB)

	assert.NotEqual(idA, idB)
}
