package services

import (
	"github.com/dong-tran/cache-play/model"
	"gorm.io/gorm"
)

type dbDateService struct {
	db *gorm.DB
}

type DbDateService interface {
	GetCurrent() *model.DbDate
}

func NewDbDateService(db *gorm.DB) DbDateService {
	return &dbDateService{
		db: db,
	}
}

func (s *dbDateService) GetCurrent() *model.DbDate {
	result := model.DbDate{}
	s.db.Raw("select current_timestamp").Scan(&result)
	return &result
}
