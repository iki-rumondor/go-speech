package repositories

import (
	"fmt"

	"github.com/iki-rumondor/go-speech/internal/domain/layers/interfaces"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MasterRepo struct {
	db *gorm.DB
}

func NewMasterInterface(db *gorm.DB) interfaces.MasterInterface {
	return &MasterRepo{
		db: db,
	}
}

func (r *MasterRepo) Create(model interface{}) error {
	return r.db.Create(model).Error
}

func (r *MasterRepo) Find(dest interface{}, condition, order string) error {
	return r.db.Preload(clause.Associations).Order(order).Find(dest, condition).Error
}

func (r *MasterRepo) First(dest interface{}, condition string) error {
	return r.db.Preload(clause.Associations).First(dest, condition).Error
}

func (r *MasterRepo) Update(model interface{}, condition string) error {
	return r.db.Where(condition).Updates(model).Error
}

func (r *MasterRepo) Delete(data interface{}, withAssociation []string) error {
	return r.db.Select(withAssociation).Delete(data).Error
}

func (r *MasterRepo) Distinct(model interface{}, column, condition string, dest *[]string) error {
	return r.db.Model(model).Distinct().Where(condition).Pluck(column, dest).Error
}

func (r *MasterRepo) Truncate(tableName string) error {
	return r.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
}