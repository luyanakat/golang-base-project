package repository

import "gorm.io/gorm"

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{
		db: db,
	}
}

func (r *LogRepository) LogInfo(message string) {
	// Save log info to database
}
