package repository

import (
	"gorm.io/gorm"
)

type groupRepositoryDb struct {
	db *gorm.DB
}

func NewGroupRepositoryDb(db *gorm.DB) groupRepositoryDb {
	return groupRepositoryDb{db: db}
}

func (r groupRepositoryDb) GetAll() ([]Group, error) {
	var groups []Group
	if err := r.db.Select("id", "name").Find(&groups); err.Error != nil {
		return nil, err.Error
	}
	return groups, nil
}

func (r groupRepositoryDb) GetById(id uint) (*Group, error) {
	var group = new(Group)
	if result := r.db.Select("id", "name").First(&group, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}
