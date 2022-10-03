package repository

import (
	"gorm.io/gorm"
	"time"
)

type accessoryStateRepository struct {
	db *gorm.DB
}

func NewAccessoryStateRepository(db *gorm.DB) accessoryStateRepository {
	return accessoryStateRepository{db: db}
}

func (r accessoryStateRepository) GetAll() ([]AccessoryState, error) {
	var accessoryStates []AccessoryState
	if err := r.db.Preload("Accessory").Preload("Group").Preload("Member").Find(&accessoryStates); err.Error != nil {
		return nil, err.Error
	}
	return accessoryStates, nil
}

func (r accessoryStateRepository) GetById(id uint64) (*AccessoryState, error) {
	var accessoryState = new(AccessoryState)
	if result := r.db.Preload("Accessory").Preload("Group").Preload("Member").First(&accessoryState, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return accessoryState, nil
}

func (r accessoryStateRepository) GetAllInGroup(id uint64) ([]AccessoryState, error) {
	var accessoryStates []AccessoryState
	if err := r.db.Preload("Accessory").Preload("Group").Preload("Member").Find(&accessoryStates, "group_id = ?", id); err.Error != nil {
		return nil, err.Error
	}
	return accessoryStates, nil
}

func (r accessoryStateRepository) GetAllInGroupByMember(groupId, memberId uint64) ([]AccessoryState, error) {
	var accessoryStates []AccessoryState
	if err := r.db.Preload("Accessory").Preload("Group").Preload("Member").Find(&accessoryStates, "group_id = ? AND member_id = ?", groupId, memberId); err.Error != nil {
		return nil, err.Error
	}
	return accessoryStates, nil
}

func (r accessoryStateRepository) GetAllByAccessoryIdInGroup(accessoryId, groupId uint64) ([]AccessoryState, error) {
	var accessoryStates []AccessoryState
	if err := r.db.Preload("Accessory").Preload("Group").Preload("Member").Find(&accessoryStates, "accessory_id = ? AND group_id = ?", accessoryId, groupId); err.Error != nil {
		return nil, err.Error
	}
	return accessoryStates, nil
}

func (r accessoryStateRepository) GetAllByAccessoryIdAndMemberIdInGroup(accessoryId, groupId, memberId uint64) ([]AccessoryState, error) {
	var accessoryStates []AccessoryState
	if err := r.db.Preload("Accessory").Preload("Group").Preload("Member").Find(&accessoryStates, "accessory_id = ? AND group_id = ? AND member_id = ?", accessoryId, groupId, memberId); err.Error != nil {
		return nil, err.Error
	}
	return accessoryStates, nil
}

func (r accessoryStateRepository) CreateAccessoryState(accessoryId, groupId, memberId uint64, state string) error {
	var accessoryState = AccessoryState{
		AccessoryId: accessoryId,
		GroupId:     groupId,
		MemberId:    memberId,
		State:       state,
		CreatedAt:   time.Now(),
	}
	if err := r.db.Create(&accessoryState); err.Error != nil {
		return err.Error
	}
	return nil
}
