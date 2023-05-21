package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port     int
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

// func LoadConfig() (*Config, error) {
// 	portStr := os.Getenv("PORT")
// 	port := 8080 // Default port
// 	if portStr != "" {
// 		var err error
// 		port, err = strconv.Atoi(portStr)
// 		if err != nil {
// 			return nil, fmt.Errorf("invalid PORT value: %s", portStr)
// 		}
// 	}

// 	dbPortStr := os.Getenv("DB_PORT")
// 	dbPort := 3306 // Default DB port
// 	if dbPortStr != "" {
// 		var err error
// 		dbPort, err = strconv.Atoi(dbPortStr)
// 		if err != nil {
// 			return nil, fmt.Errorf("invalid DB_PORT value: %s", dbPortStr)
// 		}
// 	}

// 	config := &Config{
// 		// ...
// 		Database: DatabaseConfig{
// 			Host:     os.Getenv("DB_HOST"),
// 			Port:     dbPort,
// 			Username: "username",        // Replace with your MySQL username
// 			Password: "Giridharan1729!", // Replace with your MySQL password
// 			Name:     os.Getenv("DB_NAME"),
// 		},
// 	}

// 	return config, nil
// }

func LoadConfig() (*Config, error) {
	portStr := os.Getenv("PORT")
	port := 8080 // Default port
	if portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid PORT value: %s", portStr)
		}
	}

	dbPortStr := os.Getenv("DB_PORT")
	dbPort := 3306 // Default DB port
	if dbPortStr != "" {
		var err error
		dbPort, err = strconv.Atoi(dbPortStr)
		if err != nil {
			return nil, fmt.Errorf("invalid DB_PORT value: %s", dbPortStr)
		}
	}

	config := &Config{
		Port: port,
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			Username: "username",            // Replace with your MySQL username
			Password: "Giridharan1729!",     // Replace with your MySQL password
			Name:     "home_energy_monitor", // Replace with your MySQL database name
		},
	}

	return config, nil
}
