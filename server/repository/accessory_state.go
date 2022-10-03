package repository

import "time"

type AccessoryState struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	Accessory   Accessory `gorm:"foreignKey:AccessoryId"`
	AccessoryId uint64    `json:"accessory_id"  gorm:"not null"`
	State       string    `json:"state" gorm:"type:varchar(255);not null"`
	Group       Group     `gorm:"foreignKey:GroupId"`
	GroupId     uint64    `json:"group_id" gorm:"not null"`
	Member      Member    `gorm:"foreignKey:MemberId"`
	MemberId    uint64    `json:"member_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
}

type AccessoryStateRepository interface {
	GetAll() ([]AccessoryState, error)
	GetById(uint64) (*AccessoryState, error)
	GetAllInGroup(uint64) ([]AccessoryState, error)
	GetAllInGroupByMember(uint64, uint64) ([]AccessoryState, error)
	GetAllByAccessoryIdInGroup(uint64, uint64) ([]AccessoryState, error)
	GetAllByAccessoryIdAndMemberIdInGroup(uint64, uint64, uint64) ([]AccessoryState, error)
	CreateAccessoryState(accessoryId, groupId, memberId uint64, state string) error
}
