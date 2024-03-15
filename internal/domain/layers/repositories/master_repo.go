package repositories

import "gorm.io/gorm"

type MasterRepo struct {
	db *gorm.DB
}

func NewMasterInterface(db *gorm.DB) *MasterRepo {
	return &MasterRepo{
		db: db,
	}
}
