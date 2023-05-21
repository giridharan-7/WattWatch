package services

import (
	"errors"
	"time"

	"github.com/giridharan-7/home-energy-monitor/internal/models"
	"github.com/giridharan-7/home-energy-monitor/internal/repositories"
)

type EnergyService struct {
	energyRepo *repositories.EnergyRepository
}

func NewEnergyService(energyRepo *repositories.EnergyRepository) *EnergyService {
	return &EnergyService{
		energyRepo: energyRepo,
	}
}

func (s *EnergyService) AddEnergyConsumptionData(timestamp time.Time, value float64, deviceID int, deviceName string) (*models.EnergyConsumption, error) {
	energyData := models.NewEnergyConsumption(timestamp, value, deviceID, deviceName)

	// Perform any additional validation or business logic checks before saving the data
	// ...

	// Save the energy consumption data to the repository
	err := s.energyRepo.SaveEnergyConsumptionData(energyData)
	if err != nil {
		return nil, errors.New("failed to save energy consumption data")
	}

	return energyData, nil
}

func (s *EnergyService) GetEnergyConsumptionData() ([]*models.EnergyConsumption, error) {
	data, err := s.energyRepo.GetEnergyConsumptionData()
	if err != nil {
		return nil, errors.New("failed to retrieve energy consumption data")
	}

	return data, nil
}
