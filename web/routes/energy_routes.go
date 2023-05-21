package routes

import (
	"encoding/json"
	"net/http"

	"github.com/giridharan-7/home-energy-monitor/internal/handlers"
	"github.com/gorilla/mux"
)

func SetupEnergyRoutes(router *mux.Router, energyHandler *handlers.EnergyHandler) {
	router.HandleFunc("/energy", energyHandler.GetEnergyConsumptionData).Methods(http.MethodGet)
	router.PathPrefix("/graph").Handler(http.StripPrefix("/graph", http.FileServer(http.Dir("/home/giridharan/godev/home-energy-monitor/web/"))))
	router.HandleFunc("/average-power", CalculateAveragePowerUsageHandler(energyHandler))
	router.HandleFunc("/usage-history", CalculateTotalEnergyUsageHandler(energyHandler)).Methods(http.MethodGet)
	router.HandleFunc("/total-cost", CalculateTotalCostHandle(energyHandler)).Methods(http.MethodGet)
	router.HandleFunc("/energy-goal", CalculateEnergyGoal(energyHandler)).Methods(http.MethodGet)
	router.HandleFunc("/cost-prediction", CostPrediction(energyHandler))
}

func CalculateAveragePowerUsageHandler(energyHandler *handlers.EnergyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		average := energyHandler.CalculateAveragePowerUsage()
		// Convert the average to JSON and send it as the response
		// Example:
		response := struct {
			Average float64 `json:"Average-Power"`
		}{
			Average: average,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func CostPrediction(energyHandler *handlers.EnergyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cost := (energyHandler.CalculateAveragePowerUsage() / 5) * 30
		// Convert the average to JSON and send it as the response
		// Example:
		response := struct {
			Cost float64 `json:"Predicted-Cost-For-Next-30-Days"`
		}{
			Cost: cost,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func CalculateTotalEnergyUsageHandler(energyHandler *handlers.EnergyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		total := energyHandler.CalculateTotalEnergyUsage()
		// Convert the total to JSON and send it as the response
		// Example:
		response := struct {
			Total float64 `json:"Total-Power"`
		}{
			Total: total,
			// You can use a similar approach to retrieve the total energy usage from your handler
			// based on the logic in your `retrieveEnergyData` function.
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func CalculateTotalCostHandle(energyHandler *handlers.EnergyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		total := energyHandler.CalculateTotalEnergyUsage() / 5
		// Convert the total to JSON and send it as the response
		// Example:
		response := struct {
			Total float64 `json:"Total-Cost-You-Paid"`
		}{
			Total: total,
			// You can use a similar approach to retrieve the total energy usage from your handler
			// based on the logic in your `retrieveEnergyData` function.
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func CalculateEnergyGoal(energyHandler *handlers.EnergyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lowest := energyHandler.FindLowestValue()
		average := energyHandler.CalculateAveragePowerUsage()
		// Convert the total to JSON and send it as the response
		// Example:
		response := struct {
			Total float64 `json:"Tomorrow-Goal"`
		}{
			Total: (lowest + average) / 2,
			// You can use a similar approach to retrieve the total energy usage from your handler
			// based on the logic in your `retrieveEnergyData` function.
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
