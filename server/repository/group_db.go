package repository

import (
	"gorm.io/gorm"
	"time"
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

func (r groupRepositoryDb) GetById(id uint64) (*Group, error) {
	var group = new(Group)
	if result := r.db.Select("id", "name").First(&group, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}

func (r groupRepositoryDb) CreateGroup(name, password string, time time.Time) error {
	group := Group{
		Name:      name,
		Password:  password,
		CreatedAt: time,
	}
	if result := r.db.Create(&group); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r groupRepositoryDb) CheckDuplicateName(name string) (bool, error) {
	var group = new(Group)
	if result := r.db.Select("name").First(&group, "name = ?", name); result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (r groupRepositoryDb) GetByIdWithPassword(id uint64) (*Group, error) {
	var group = new(Group)
	if result := r.db.Select("id", "name", "password").First(&group, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}
