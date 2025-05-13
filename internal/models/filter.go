package models

import "time"

type Filter struct {
	ID                   []uint
	StudentCode          []string
	FirstName            []string
	LastName             []string
	MiddleName           []string
	Email                []string
	Sign                 []string
	Phone                []string
	Gender               *uint
	DOB                  []time.Time
	Address              []string
	Avatar               []string
	IsActive             *bool
	RoomID               []string
	NotificationChannels []uint
}
