package model

type Role struct {
	ID   string `gorm:"type:char(6)"`
	Name string `gorm:"type:enum('student', 'manager', 'admin'); index"`
}
