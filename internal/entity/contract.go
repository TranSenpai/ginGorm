package entity

import (
	"time"
)

// type User interface {
// 	RegistryRoom()
// }

// varchar in MySQL use 1 byte (2^8 - 1 = 255) to store length of record and another one - four byte each character
// the reason why subtract 1 for store the & to connect the string
// so if the record's length is < 255, we should declare the length is 255 to take the most of capacity of bit length
// and if the record's length is > 255 and < 510 use the size 510 -> length = 2^n - 1
type Contract struct {
	StudentCode          string     `gorm:"type:char(10);"`
	FullName             string     `gorm:"type:varchar(255);"`
	Email                string     `gorm:"type:varchar(255);"`
	Sign                 string     `gorm:"type:varchar(255);"` // Because bcrypt algorithm return 60 characteres
	Phone                string     `gorm:"type:char(10);"`
	Gender               uint8      `gorm:"type:tinyint UNSIGNED; default: 1"`
	DOB                  *time.Time `gorm:"type:date"`
	Address              *string    `gorm:"type:varchar(255)"`
	Avatar               []byte     `gorm:"type:mediumblob"` // Should store varchar(510)
	IsActive             bool       `gorm:"type:bool"`
	RegistryAt           *time.Time `gorm:"type:timestamp; autoCreateTime;"`
	LoginAt              *time.Time `gorm:"type:timestamp; autoUpdateTime"`
	RoomID               string     `gorm:"type:char(5)"`
	NotificationChannels uint8      `gorm:"type:tinyint UNSIGNED; default: 1"`
}
