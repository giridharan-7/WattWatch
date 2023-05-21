package database

import (
	"database/sql"
	"fmt"

	"github.com/giridharan-7/home-energy-monitor/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(cfg *config.DatabaseConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type EnergyRepository struct {
	db *sql.DB
}

func NewEnergyRepository(db *sql.DB) *EnergyRepository {
	return &EnergyRepository{
		db: db,
	}
}
