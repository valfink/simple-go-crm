package customer

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

func PrepareMockData() map[uuid.UUID]Customer {
	slog.Info("Preparing mock data...")
	mockCustomers := make(map[uuid.UUID]Customer, 100)
	for i := 0; i < 100; i++ {
		newId, err := uuid.NewRandom()
		if err != nil {
			slog.Error("Could not generate uuid", "Error", err)
			continue
		}

		mockCustomers[newId] = Customer{
			ID:        newId,
			Name:      fmt.Sprintf("Fake Customer %d", i),
			Role:      fmt.Sprintf("Company Role %d", i),
			Email:     fmt.Sprintf("customer%d@gmail.com", i),
			Phone:     fmt.Sprintf("+%d 123 456", i),
			Contacted: i%2 == 0,
		}
	}

	return mockCustomers
}
