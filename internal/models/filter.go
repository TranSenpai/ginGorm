package models

import "time"

type Filter struct {
	StudentCode          []string
	FirstName            *string
	LastName             *string
	MiddleName           *string
	Email                []string
	Sign                 []string
	Phone                []string
	Gender               []uint8
	DOB                  []time.Time
	Address              []string
	Avatar               []string
	IsActive             *bool
	RoomID               []string
	NotificationChannels []uint
}
