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
	StudentCode          string     `json:"student_code,omitempty" binding:"required"`
	FirstName            *string    `json:"first_name,omitempty"`
	LastName             *string    `json:"last_name,omitempty"`
	MiddleName           *string    `json:"middle_name,omitempty"`
	Email                string     `json:"email,omitempty" binding:"required"`
	Sign                 string     `json:"sign,omitempty" binding:"required"`
	Phone                string     `json:"phone,omitempty" binding:"required"`
	Gender               *uint8     `json:"gender,omitempty"`
	DOB                  *time.Time `json:"dob,omitempty"`
	Address              *string    `json:"address,omitempty"`
	Avatar               *string    `json:"avatar,omitempty"`
	IsActive             *bool      `json:"is_active,omitempty"`
	RegistryAt           *time.Time `json:"registry_at,omitempty"`
	LoginAt              *time.Time `json:"login_at,omitempty"`
	RoomID               *string    `json:"room_id,omitempty"`
	NotificationChannels *uint8     `json:"notification_channels,omitempty"`
}
