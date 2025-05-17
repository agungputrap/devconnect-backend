package postgres

import (
	"github.com/agungputrap/devconnect-backend/internal/infrastructure/config"
	"gorm.io/gorm"
)

type DatabaseProvider struct {
	db *gorm.DB
}

func NewDatabaseProvider(cfg *config.DatabaseConfig) (*DatabaseProvider, error) {
	database, err := config.NewDatabase(cfg)
	if err != nil {
		return nil, err
	}

	return &DatabaseProvider{db: database.DB()}, nil
}

func (p *DatabaseProvider) DB() *gorm.DB {
	return p.db
}
