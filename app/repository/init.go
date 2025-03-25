package repository

import "gorm.io/gorm"

type BaseRepository struct {
	LogRepo *LogRepository
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		LogRepo: NewLogRepository(db),
	}
}
