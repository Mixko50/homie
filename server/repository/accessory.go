package repository

import "time"

type Accessory struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Group     Group     `gorm:"foreignKey:GroupId"`
	GroupId   uint64    `json:"group_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type AccessoryRepository interface {
	GetAll() ([]Accessory, error)
	GetById(uint64) (*Accessory, error)
	GetAllInGroup(uint64) ([]Accessory, error)
	CreateAccessory(name string, groupId uint64, time time.Time) error
	CheckDuplicateNameInGroup(name string, id uint64) (bool, error)
}
