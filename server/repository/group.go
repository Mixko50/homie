package repository

type Group struct {
	Id       uint64 `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

type GroupRepository interface {
	GetAll() ([]Group, error)
	GetById(uint) (*Group, error)
}
