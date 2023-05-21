package repositories

import (
	"sync"

	"github.com/giridharan-7/home-energy-monitor/internal/models"
)

type EnergyRepository struct {
	dataMutex sync.Mutex
	data      []*models.EnergyConsumption
}

func NewEnergyRepository() *EnergyRepository {
	return &EnergyRepository{
		data: []*models.EnergyConsumption{},
	}
}

func (r *EnergyRepository) SaveEnergyConsumptionData(data *models.EnergyConsumption) error {
	r.dataMutex.Lock()
	defer r.dataMutex.Unlock()

	// Assign an ID to the data
	data.ID = len(r.data) + 1

	// Add the data to the repository
	r.data = append(r.data, data)

	return nil
}

func (r *EnergyRepository) GetEnergyConsumptionData() ([]*models.EnergyConsumption, error) {
	r.dataMutex.Lock()
	defer r.dataMutex.Unlock()

	// Return a copy of the data in the repository
	dataCopy := make([]*models.EnergyConsumption, len(r.data))
	copy(dataCopy, r.data)

	return dataCopy, nil
}
