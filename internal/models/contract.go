package models

import (
	"time"
)

// In Golang package time Time
// The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
// As this time is unlikely to come up in practice, the Time.IsZero method
// gives a simple way of detecting a time that has not been initialized explicitly.

// In MySQL
// The NO_ZERO_IN_DATE mode affects whether the server permits dates in
// which the year part is nonzero but the month or day part is 0.
// (This mode affects dates such as '2010-00-01' or '2010-01-00',
// but not '0000-00-00'. To control whether the server permits '0000-00-00',
// use the NO_ZERO_DATE mode.) The effect of NO_ZERO_IN_DATE also depends on
// whether strict SQL mode is enabled.

type Contract struct {
	StudentCode          string     `json:"StudentCode,omitempty"`
	FirstName            *string    `json:"FirstName,omitempty"`
	LastName             *string    `json:"LastName,omitempty"`
	MiddleName           *string    `json:"MiddleName,omitempty"`
	Email                string     `json:"Email,omitempty"`
	Sign                 string     `json:"Sign,omitempty"`
	Phone                string     `json:"Phone,omitempty"`
	Gender               *uint8     `json:"Gender,omitempty"`
	DOB                  *time.Time `json:"DOB,omitempty"`
	Address              *string    `json:"Address,omitempty"`
	Avatar               *string    `json:"Avatar,omitempty"`
	IsActive             *bool      `json:"IsActive,omitempty"`
	RegistryAt           *time.Time `json:"RegistryAt,omitempty"`
	LoginAt              *time.Time `json:"LoginAt,omitempty"`
	RoomID               *string    `json:"RoomID,omitempty"`
	NotificationChannels *uint8     `json:"NotificationChannels,omitempty"`
}
