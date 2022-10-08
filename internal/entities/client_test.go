package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientWithValidData(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, 0, client.Points)
}

func TestNewClientWithInvalidDate(t *testing.T) {
	client, err := NewClient("", "j@j.com")

	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client name is required")

	client, err = NewClient("Jailton", "")

	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client e-mail is required")
}

func TestAddPointsBatch(t *testing.T) {
	pointsTable := []int{2, 4, 6, 8, 10}

	for _, point := range pointsTable {
		client, _ := NewClient("Jailton", "j@j")
		err := client.AddPoints(point)

		assert.Nil(t, err)
		if client.Points != point {
			t.Errorf("Points expected: %d, got: %d", point, client.Points)
		}
	}
}

func FuzzClient_AddPoints(f *testing.F) {
	seeding := []int{2, 4, 6, 8, 10}

	for _, seed := range seeding {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, point int) {
		client, _ := NewClient("Jailton", "j@j")
		err := client.AddPoints(point)

		if err != nil {
			return
		}

		if client.Points != point {
			t.Errorf("Points expected: %d, got: %d", point, client.Points)
		}
	})
}
