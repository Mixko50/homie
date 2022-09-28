package repository

import "time"

type Member struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	Group     Group     `gorm:"foreignKey:GroupId"`
	GroupId   uint64    `json:"group_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	UserAgent string    `json:"user_agent" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type MemberRepository interface {
	GetAll() ([]Member, error)
	GetById(uint64) (*Member, error)
	CreateMember(name string, groupId uint64, agent string, time time.Time) (*Member, error)
	CheckDuplicateNameInMember(name string, id uint64) (bool, error)
}
