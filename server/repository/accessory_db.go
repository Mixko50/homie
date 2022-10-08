package repository

import (
	"gorm.io/gorm"
	"time"
)

type accessoryRepositoryDB struct {
	db *gorm.DB
}

func NewAccessoryRepositoryDB(db *gorm.DB) accessoryRepositoryDB {
	return accessoryRepositoryDB{db: db}
}

func (r accessoryRepositoryDB) GetAll() ([]Accessory, error) {
	var accessories []Accessory
	if err := r.db.Preload("Group").Select("id", "name", "group_id", "created_at", "updated_at").Find(&accessories); err.Error != nil {
		return nil, err.Error
	}
	return accessories, nil
}

func (r accessoryRepositoryDB) GetById(id uint64) (*Accessory, error) {
	var accessory = new(Accessory)
	if result := r.db.Preload("Group").Select("id", "name", "group_id", "created_at", "updated_at").First(&accessory, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return accessory, nil
}

func (r accessoryRepositoryDB) CreateAccessory(name string, groupId uint64, time time.Time) error {
	accessory := Accessory{
		Name:      name,
		GroupId:   groupId,
		CreatedAt: time,
		UpdatedAt: time,
	}
	if result := r.db.Create(&accessory); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r accessoryRepositoryDB) CheckDuplicateNameInGroup(name string, id uint64) (bool, error) {
	var accessory = new(Accessory)
	if result := r.db.Preload("Group").Select("name").First(&accessory, "name = ? AND group_id = ?", name, id); result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (r accessoryRepositoryDB) GetAllInGroup(groupId uint64) ([]Accessory, error) {
	var accessories []Accessory
	if err := r.db.Preload("Group").Select("id", "name", "group_id", "created_at", "updated_at").Where("group_id = ?", groupId).Find(&accessories); err.Error != nil {
		return nil, err.Error
	}
	return accessories, nil
}
