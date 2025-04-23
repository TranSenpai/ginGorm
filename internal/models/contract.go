package model

import (
	"time"
)

// type User interface {
// 	RegistryRoom()
// }

type Contract struct {
	ID          string    `gorm:"primaryKey;type:char(12)"`
	StudentCode string    `gorm:"type:char(10); index"`
	FullName    string    `gorm:"type:varchar(100); index"`
	Email       string    `gorm:"type:varchar(100); index"`
	Sign        string    `gorm:"type:varchar(100); not null"` // Vì bcrypt trả về chuỗi 60 ký tự
	Phone       string    `gorm:"type:char(10); index"`
	Gender      string    `gorm:"type:enum('male', 'female', 'other'); default:'male'"` // thường tạo ngoài, ít tạo trong db, enum thường là số
	DOB         time.Time `gorm:"type:date"`
	Address     string    `gorm:"type:text"`
	Avatar      []byte    `gorm:"type:mediumblob"` // ảnh avatar thường ở 500KB – 1MB (500,000 – 1,000,000 bytes)
	IsActive    bool      `gorm:"type:bool"`
	RegistryAt  time.Time `gorm:"type:timestamp; autoCreateTime"`
	LoginAt     time.Time `gorm:"type:timestamp; autoUpdateTime"`
	RoomID      string    `gorm:"type:char(5)"`
	// Room                 Room      `gorm:"foreignKey:RoomID; references:ID"` // belong to 1 Room
	NotificationChannels string `gorm:"type:set('email','sms','zalo')"` // Như enum nhưng cho phép bỏ nhiều giá trị trong set
	// autoUpdateTime là directive do GORM xử lý ở tầng ORM,
	// nghĩa là: Khi call .Save(), .Updates(), .Update()... GORM
	// sẽ tự động gán giá trị thời gian hiện tại (time.Now()) cho trường đó.
}
