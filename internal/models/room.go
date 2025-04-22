package model

import (
	"gorm.io/datatypes"
)

type Room struct {
	ID                string         `gorm:"type:char(5);primaryKey"` // ID sẽ có cấu trúc 1-2-2 : Building-Floor-RoomID
	RoomNumber        uint16         `gorm:"type:smallint unsigned"`
	Building          string         `gorm:"type:char(1)"`
	Floor             uint8          `gorm:"type:tinyint"`
	Capacity          uint8          `gorm:"type:tinyint"`
	CurrentOccupants  uint8          `gorm:"type:tinyint; "` //check:current_occupants <= capacity chậm db, check bằng code
	HasAirconditioner bool           `gorm:"type:bool"`
	MonthlyFee        float64        `gorm:"type:decimal(10,2)"`
	Status            string         `gorm:"type:enum('available', 'maintenance', 'full'); default:'available'"`
	Description       string         `gorm:"type:text"`
	Facilities        datatypes.JSON `gorm:"type:json"`
	// Contracts         []Contract       `gorm:"foreignKey:RoomID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // has many user
}
