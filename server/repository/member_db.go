package repository

import (
	"gorm.io/gorm"
	"time"
)

type memberRepositoryDb struct {
	db *gorm.DB
}

func NewMemberRepositoryDb(db *gorm.DB) memberRepositoryDb {
	return memberRepositoryDb{db: db}
}

func (r memberRepositoryDb) GetAll() ([]Member, error) {
	var members []Member
	if err := r.db.Select("id", "name", "group_id", "created_at", "updated_at").Find(&members); err.Error != nil {
		return nil, err.Error
	}
	return members, nil
}

func (r memberRepositoryDb) GetById(id uint64) (*Member, error) {
	var member Member
	if err := r.db.Select("id", "name", "group_id", "created_at", "updated_at").Where("id = ?", id).First(&member); err.Error != nil {
		return nil, err.Error
	}
	return &member, nil
}

func (r memberRepositoryDb) CreateMember(name, password string, groupId uint64, agent string, time time.Time) (*Member, error) {
	member := Member{
		Name:      name,
		GroupId:   groupId,
		Password:  password,
		UserAgent: agent,
		CreatedAt: time,
		UpdatedAt: time,
	}
	if result := r.db.Create(&member); result.Error != nil {
		return nil, result.Error
	}
	return &member, nil
}

func (r memberRepositoryDb) CheckDuplicateNameInMember(name string, id uint64) (bool, error) {
	var member = new(Member)
	if result := r.db.Select("name").First(&member, "name = ? AND group_id = ?", name, id); result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (r memberRepositoryDb) GetMembersByName(name string) ([]Member, error) {
	var members []Member
	if result := r.db.Select("id", "name", "group_id", "password").First(&members, "name = ?", name); result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, result.Error
		}
	}
	return members, nil
}
