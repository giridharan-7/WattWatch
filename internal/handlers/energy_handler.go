package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type EnergyHandler struct {
	// Add any dependencies or services required by the energy handler
}

func NewEnergyHandler() *EnergyHandler {
	// Initialize and return a new instance of the energy handler
	return &EnergyHandler{}
}

func (h *EnergyHandler) retrieveEnergyData() ([]EnergyConsumption, error) {
	// Implement the logic to retrieve energy consumption data from a service or repository
	// For example, you can call a function from the energy service or repository
	// and return the retrieved data as a slice of EnergyConsumption structs
	// Example:
	return []EnergyConsumption{
		{ID: 1, Timestamp: "21-5-23", Units: 6},
		{ID: 2, Timestamp: "22-5-23", Units: 7},
		{ID: 4, Timestamp: "24-5-23", Units: 8},
		{ID: 3, Timestamp: "23-5-23", Units: 5},
		{ID: 5, Timestamp: "25-5-23", Units: 6},
		{ID: 6, Timestamp: "26-5-23", Units: 5},
		{ID: 7, Timestamp: "27-5-23", Units: 6},
		{ID: 8, Timestamp: "28-5-23", Units: 7},
		{ID: 9, Timestamp: "29-5-23", Units: 8},
		{ID: 10, Timestamp: "30-5-23", Units: 9},
		{ID: 11, Timestamp: "31-5-23", Units: 5},
		{ID: 12, Timestamp: "01-6-23", Units: 7},
		{ID: 13, Timestamp: "02-6-23", Units: 3},
		{ID: 14, Timestamp: "03-6-23", Units: 5},
		{ID: 15, Timestamp: "04-6-23", Units: 4},
	}, nil
}

func (h *EnergyHandler) GetEnergyConsumptionData(w http.ResponseWriter, r *http.Request) {
	// Handle GET request to retrieve energy consumption data
	// Retrieve energy consumption data from the service or repository
	data, err := h.retrieveEnergyData()
	if err != nil {
		log.Printf("Failed to retrieve energy consumption data: %v\n", err)
		http.Error(w, "Failed to retrieve energy consumption data", http.StatusInternalServerError)
		return
	}
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Send the response back to the client
	json.NewEncoder(w).Encode(data)

}

func (h *EnergyHandler) CalculateAveragePowerUsage() float64 {
	data, err := h.retrieveEnergyData()
	if err != nil {
		// Handle the error
		return 0.0
	}

	total := 0.0
	for _, d := range data {
		total += float64(d.Units)
	}

	average := total / float64(len(data))
	return average
}

func (h *EnergyHandler) CalculateTotalEnergyUsage() float64 {
	data, err := h.retrieveEnergyData()
	if err != nil {
		return -0.1
	}
	total := 0.0
	for _, d := range data {
		total += d.Units
	}

	return total
}

func (h *EnergyHandler) FindLowestValue() float64 {
	data, err := h.retrieveEnergyData()
	if err != nil {
		return 0.0
	}

	if len(data) == 0 {
		return 0.0
	}

	lowest := data[0].Units
	for _, d := range data {
		if d.Units < lowest {
			lowest = d.Units
		}
	}

	return lowest
}

func (h *EnergyHandler) SaveEnergyConsumptionData(w http.ResponseWriter, r *http.Request) {
	// Handle POST request to save energy consumption data

	// Parse the request body and extract the energy consumption data
	var data EnergyConsumption
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Failed to parse request body: %v\n", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Save the energy consumption data to the service or repository
	err = h.saveEnergyData(data)
	if err != nil {
		log.Printf("Failed to save energy consumption data: %v\n", err)
		http.Error(w, "Failed to save energy consumption data", http.StatusInternalServerError)
		return
	}

	// Send a success response back to the client
	w.WriteHeader(http.StatusCreated)
}

func (h *EnergyHandler) saveEnergyData(data EnergyConsumption) error {
	// Implement the logic to save energy consumption data to a service or repository
	// For example, you can call a function from the energy service or repository
	// and pass the data to be saved
	// Example:
	log.Printf("Saving energy consumption data: %+v\n", data)
	// ...

	return nil
}

type EnergyConsumption struct {
	ID        int     `json:"id"`
	Timestamp string  `json:"timestamp"`
	Units     float64 `json:"units"`
}
