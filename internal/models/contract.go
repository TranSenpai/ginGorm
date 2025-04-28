package models

import (
	"time"
)

type Contract struct {
	StudentCode          string
	FullName             string
	Email                string
	Sign                 string
	Phone                string
	Gender               uint8
	DOB                  time.Time
	Address              string
	Avatar               string
	IsActive             bool
	RegistryAt           time.Time
	LoginAt              time.Time
	RoomID               string
	NotificationChannels uint8
}
